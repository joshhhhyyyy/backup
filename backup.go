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
	key := flag.String("key", os.Getenv("KEY"), "the api key for this sentry project (required)")
	message := flag.String("m", time.Now().Format("🌈 02 Jan"), "the commit message (not required)")
	bup := flag.String("bup", "joseos.com", "the betteruptime heartbeat to GET (optional)")
	flag.Parse()

	log.Println("using sentry key:", *key)

	uuuuuuuuu := sentry.Init(sentry.ClientOptions{
		Dsn: *key,
		TracesSampleRate: 1.0,
	})
	if uuuuuuuuu != nil {
		log.Fatalf("sentry.Init: %s", uuuuuuuuu)
		panic(uuuuuuuuu)
	}

	log.Println("commit message: ", *message)

	gitpull, pullerr := exec.Command("git", "pull").Output()
	log.Println(string(gitpull))
	if pullerr != nil {
		sentry.CaptureMessage(string(gitpull))
		log.Println("there was an error when performing git add")
		panic(pullerr)
	}

	// check if the length of a variable is zero and returns true if it is zero.
	gitstatus, gitstatuserr := exec.Command("git", "status", "--porcelain").Output()
	if gitstatuserr != nil {
		panic(gitstatuserr)
	}
	if len(string(gitstatus)) == 0 {
		log.Printf("There are no changes to be committed.")
		os.Exit(0)
	}

	gitadd, adderr := exec.Command("git", "add", ".").Output()
	log.Println(string(gitadd))
	if adderr != nil {
		sentry.CaptureMessage(string(gitadd))
		log.Println("there was an error when performing git add")
		panic(adderr)
	}

	gitcommit, commiterr := exec.Command("git", "commit", "-m", *message).Output()
	log.Println(string(gitcommit))
	if commiterr != nil {
		sentry.CaptureMessage(string(gitcommit))
		log.Println("there was an error when performing git commit")
		panic(commiterr)
	}

	gitpush, pusherr := exec.Command("git", "push", "origin", "main").Output()
	log.Println(string(gitpush))
	if pusherr != nil {
		sentry.CaptureMessage(string(gitpush))
		log.Println("there was an error when performing git push!!!!")
		panic(pusherr)
	}

	if *bup != "joseos.com" {
		http.Get(*bup)
	}
}
