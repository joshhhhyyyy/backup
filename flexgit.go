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

	key := flag.String("key", os.Getenv("KEY"), "the api key for this sentry project")
	flag.Parse()

	log.Println("using sentry key:", *key)

	uuuuuuuuu := sentry.Init(sentry.ClientOptions{
		Dsn: *key,
	})
	if uuuuuuuuu != nil {
		log.Fatalf("sentry.Init: %s", uuuuuuuuu)
	}

	lmao := time.Now().Format("🌈 02 Jan")
	log.Println(lmao)

	gitpull, pullerr := exec.Command("git", "pull").Output()
	log.Println(string(gitpull))
	if pullerr != nil {
		log.Println(pullerr)
		sentry.CaptureException(pullerr)
		log.Println("there was an error when performing git add")
		erroraaa = true
	}

	gitadd, adderr := exec.Command("git", "add", ".").Output()
	log.Println(string(gitadd))
	if adderr != nil {
		log.Println(adderr)
		sentry.CaptureException(adderr)
		log.Println("there was an error when performing git add")
		erroraaa = true
	}

	gitcommit, commiterr := exec.Command("git", "commit", "-m", lmao).Output()
	log.Println(string(gitcommit))
	if commiterr != nil {
		log.Println(commiterr)
		sentry.CaptureException(commiterr)
		log.Println("there was an error when performing git commit")
		erroraaa = true
	}

	gitpush, pusherr := exec.Command("git", "push", "origin", "main").Output()
	log.Println(string(gitpush))
	if pusherr != nil {
		log.Println(pusherr)
		sentry.CaptureException(pusherr)
		log.Println("there was an error when performing git push")
		erroraaa = true
	}

	// check if boolean is true
	if !erroraaa {
	http.Get("https://betteruptime.com/api/v1/heartbeat/zfnS1uFSeYdSwQY41Na7mMRW")
	} else {
		log.Println("there was an error")
	}
}