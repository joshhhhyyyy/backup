package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	var erroraaa bool = false
	key := flag.String("key", "https://nil", "the api key for this sentry project")
	flag.Parse()

	uuuuuuuuu := sentry.Init(sentry.ClientOptions{
		Dsn: *key,
	})
	if uuuuuuuuu != nil {
		log.Fatalf("sentry.Init: %s", uuuuuuuuu)
	}

	lmao := time.Now().Format("🌈 02 Jan")

	gitadd := exec.Command("git", "add", ".").Run()
	if gitadd != nil {
		sentry.CaptureException(gitadd)
		log.Println("there was an error when performing git add")
		erroraaa = true
	}

	gitcommit := exec.Command("git", "commit", "-m", lmao).Run()
	if gitcommit != nil {
		sentry.CaptureException(gitcommit)
		log.Println("there was an error when performing git commit")
		erroraaa = true
	}

	gitpush := exec.Command("git", "push", "origin", "main").Run()
	if gitpush != nil {
		sentry.CaptureException(gitpush)
		log.Println("there was an error when performing git push")
		erroraaa = true
	}

	// check if boolean is true
	if erroraaa {
	http.Get("https://betteruptime.com/api/v1/heartbeat/zfnS1uFSeYdSwQY41Na7mMRW")
	log.Println("there was an error")
	}
}