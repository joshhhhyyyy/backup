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
	var erroraaa bool = false

	key := flag.String("key", os.Getenv("KEY"), "the api key for this sentry project (required)")
	message := flag.String("message", time.Now().Format("🌈 02 Jan"), "the commit message (not required)")
	bup := flag.String("bup", "joseos.com", "the betteruptime heartbeat to GET (optional)")
	flag.Parse()

	log.Println("using sentry key:", *key)

	uuuuuuuuu := sentry.Init(sentry.ClientOptions{
		Dsn:              *key,
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
		log.Println(pullerr)
		sentry.CaptureMessage(string(gitpull))
		log.Println("there was an error when performing git add")
		erroraaa = true
	}

	gitadd, adderr := exec.Command("git", "add", ".").Output()
	log.Println(string(gitadd))
	if adderr != nil {
		log.Println(adderr)
		sentry.CaptureMessage(string(gitadd))
		log.Println("there was an error when performing git add")
		erroraaa = true
	}

	gitcommit, commiterr := exec.Command("git", "commit", "-m", *message).Output()
	log.Println(string(gitcommit))
	if commiterr != nil {
		log.Println(commiterr)
		sentry.CaptureMessage(string(gitcommit))
		log.Println("there was an error when performing git commit")
		erroraaa = true
	}

	gitpush, pusherr := exec.Command("git", "push", "origin", "main").Output()
	log.Println(string(gitpush))
	if pusherr != nil {
		log.Println(pusherr)
		sentry.CaptureMessage(string(gitpush))
		log.Println("there was an error when performing git push!!!!")
		erroraaa = true
	}

	// check if an error has occured
	if !erroraaa && *bup != "joseos.com" {
		http.Get(*bup)
	} else if !erroraaa{
		log.Println("there was no error")
	} else if erroraaa {
		log.Println("there was an error")
	}
}
