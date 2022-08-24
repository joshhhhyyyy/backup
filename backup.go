package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Initiallise flags
	key := flag.String("key", "https://ce35494a70c94719bc09c3e1086517ad@o1153157.ingest.sentry.io/6233195", "the api key for this sentry project (optional)")
	message := flag.String("m", time.Now().Format("🌈 02 Jan"), "manual override for the commit message (not required)")
	ping := flag.String("ping", "nil", "send a http get request everytime this runs for uptime monitoring (optional)")
	flag.Parse()

	// Print stuff
	if *key != "https://ce35494a70c94719bc09c3e1086517ad@o1153157.ingest.sentry.io/6233195" {
		log.Println("using custom sentry key:", *key)
	}

	if *ping != "nil" {
		log.Println("Will ping:", *ping)
	}

	if *message != time.Now().Format("🌈 02 Jan") {
		log.Println("custom commit message: ", *message)
	}

	// Initiallise Sentry
	sentryerr := sentry.Init(sentry.ClientOptions{
		Dsn:              *key,
		TracesSampleRate: 1.0,
	})
	if sentryerr != nil {
		log.Fatalf("sentry.Init: %s", sentryerr)
		panic(sentryerr)
	}

	gitpull, pullerr := exec.Command("git", "pull").Output()
	log.Println(string(gitpull))
	if pullerr != nil {
		sentry.CaptureMessage(string(gitpull))
		log.Println("there was an error when performing git pull")
		log.Println("Continuing,")
	}

	// check if the length of git status is zero and returns true if it is zero.
	gitstatus, gitstatuserr := exec.Command("git", "status", "--porcelain").Output()
	log.Println(gitstatus)
	if gitstatuserr != nil {
		panic(gitstatuserr)
	}

	if len(string(gitstatus)) == 0 {
		log.Printf("There are no changes to be committed.")
		if *ping != "nil" {
			httpget, httperr := http.Get(*ping)
			if httperr != nil {
				log.Println("there was an error when perfoming http.get after git status.")
				sentry.CaptureException(httperr)
				log.Println(httpget)
				panic(httperr)
			}
		}
		os.Exit(0)
	}

	// Stage and commit all files
	gitcommit, commiterr := exec.Command("git", "commit", "-am", *message).Output()
	log.Println(string(gitcommit))
	if commiterr != nil {
		sentry.CaptureMessage(string(gitcommit))
		log.Println("there was an error when performing git commit")
		panic(gitcommit)
	}

	// Push all files
	gitpush, pusherr := exec.Command("git", "push").Output()
	log.Println(string(gitpush))
	if pusherr != nil {
		sentry.CaptureMessage(string(gitpush))
		log.Println("there was an error when performing git push")
		panic(pusherr)
	}

	// http.get provided link if it isnt nil
	if *ping != "nil" {
		httpget, httperr := http.Get(*ping)
		if httperr != nil {
			log.Println("there was an error when pinging the supplied url")
			sentry.CaptureException(httperr)
			log.Println(httpget)
			panic(httperr)
		}
	}
}
