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
	bup := flag.String("bup", "nil", "a cronjob tracker http link  (eg. betteruptime heartbeat) to http.GET (optional)")
	flag.Parse()

	// Print stuff
	if *key != "https://ce35494a70c94719bc09c3e1086517ad@o1153157.ingest.sentry.io/6233195" {
		log.Println("using custom sentry key:", *key)
	}

	if *bup != "nil" {
		log.Println("with betteruptime key:", *bup)
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
		if *bup != "nil" {
			httpget, httperr := http.Get(*bup)
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

		// try again
		log.Println("trying again,")
		onerrgitcommit := exec.Command("git", "commit", "-am", *message)
		onerrgitcommit.Stdin = os.Stdin
		onerrgitcommit.Stdout = os.Stdout
		onerrgitcommit.Stderr = os.Stderr
		onerrgitcommiterr := onerrgitcommit.Run()

		if onerrgitcommiterr != nil {
			sentry.CaptureException(onerrgitcommiterr)
			log.Println(onerrgitcommiterr)
			panic(onerrgitcommiterr)
		}
	}

	// Push all files
	gitpush, pusherr := exec.Command("git", "push").Output()
	log.Println(string(gitpush))
	if pusherr != nil {
		sentry.CaptureMessage(string(gitpush))
		log.Println("there was an error when performing git push")

		// try again
		log.Println("trying again,")
		onerrgitpush := exec.Command("git", "push")
		onerrgitpush.Stdin = os.Stdin
		onerrgitpush.Stdout = os.Stdout
		onerrgitpush.Stderr = os.Stderr
		onerrgitpusherr := onerrgitpush.Run()

		if onerrgitpusherr != nil {
			sentry.CaptureException(onerrgitpusherr)
			log.Println(onerrgitpusherr)
			panic(onerrgitpusherr)
		}
	}
	// http.get provided link if it isnt nil
	if *bup != "nil" {
		httpget, httperr := http.Get(*bup)
		if httperr != nil {
			log.Println("there was an error when perfoming http.get at the end.")
			sentry.CaptureException(httperr)
			log.Println(httpget)
			panic(httperr)
		}
	}
}
