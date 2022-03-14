# Backup &nbsp; &nbsp; &nbsp;[![Latest Release](https://img.shields.io/github/release/joshhhhyyyy/Backup.svg)](https://github.com/joshhhhyyyy/Backup/releases)      [![Go ReportCard](https://goreportcard.com/badge/joshhhhyyyy/Backup)](https://goreportcard.com/report/joshhhhyyyy/Backup)
![Yep](https://socialify.git.ci/joshhhhyyyy/Backup/image?descriptionEditable=backup%20a%20database%20using%20git.%20Err%20handling%20by%20Sentry%2C%20Reporting%20by%20Betteruptime.%20Made%20with%20%F0%9F%A9%B8%20%2C%20%F0%9F%98%93%20%26%20%F0%9F%98%AD&font=Source%20Code%20Pro&language=1&owner=1&pattern=Overlapping%20Hexagons&theme=Dark)

## What is this?
**A Simple and Lightweight program to automatically backup a database using git.**

Err handling by **[Sentry](sentry.io)**, Uses heartbeats by **[Betteruptime](https://betteruptime.com)** 

Made with 🩸 , 😓 &amp; 😭

## Installation
### Via Go
```go get github.com/joshhhhyyyy/backup```

```go install github.com/joshhhhyyyy/backup```

```export PATH=$PATH:$(go env GOPATH)/bin``` (Add gopath to path)

### Via apt
```echo "deb [trusted=yes] https://apt.joseos.com/ ./" | sudo tee /etc/apt/sources.list.d/joseos.list```

```sudo apt update```

```sudo apt install backup```

## Usage
### Systemd Timers
[Sample Service file](https://github.com/storageroom/storage/blob/main/linux/systemd/flexgit.service) | [Sample Timer file](https://github.com/storageroom/storage/blob/main/linux/systemd/flexgit.timer)

Basically, systemd timers is a great alternative to cronjobs as it can be invoked and controlled via systemctl.

Every timer has a corresponding service file that it invokes when the time specified is hit. In the sample files provided, the timer triggers a backup at 4am and 59seconds daily.

### Manual Usage
This program can also be set with flags, where you can specify the commit message and cronjob tracker of your choice.

```backup [OPTIONS]```

## Options
**note: both single minus "-" and double minus "--" work fine

```-bup=""``` // Optional, Type: string, a cronjob tracker http link  (eg. betteruptime heartbeat) to http.GET

```-message=""``` // Not needed at all, Type: string, optional commit message to pass for manual runs

## Debug Options

```-key=""``` // Optional, Type: string, Custom [Sentry.io](sentry.io) dsn (key url) for project