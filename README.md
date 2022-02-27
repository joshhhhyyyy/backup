# backup
![Yep](https://socialify.git.ci/joshhhhyyyy/Backup/image?descriptionEditable=backup%20a%20database%20using%20git.%20Err%20handling%20by%20Sentry%2C%20Reporting%20by%20Betteruptime.%20Made%20with%20%F0%9F%A9%B8%20%2C%20%F0%9F%98%93%20%26%20%F0%9F%98%AD&font=Source%20Code%20Pro&language=1&owner=1&pattern=Overlapping%20Hexagons&theme=Dark)

## What is this?
**A Simple program to automatically backup a database using git. Err handling by Sentry, Reporting by Betteruptime. Made with 🩸 , 😓 &amp; 😭**

## Installation
```go get github.com/joshhhhyyyy/backup```

```go install github.com/joshhhhyyyy/backup```

```export PATH=$PATH:$(go env GOPATH)/bin``` (Add gopath to path)

## Usage
### Systemd Timers
[Sample Service file](https://github.com/storageroom/storage/blob/main/linux/systemd/flexgit.service)

[Sample Timer file](https://github.com/storageroom/storage/blob/main/linux/systemd/flexgit.service)

Basically, systemd timers is a great alternative to cronjobs as it can be invoked and controlled via systemctl.

Every timer has a corresponding service file that it invokes when the time specified is hit. In the sample files provided, the timer triggers a backup at 4am and 59seconds daily.

### Manual Usage
This program can also be run manually with the command, where you can specify the commit message as well

```backup [OPTIONS]```

## Options
**note: both single minus "-" and double minus "--" work fine

```-key=""``` // Required, Type: string, [Sentry.io](sentry.io) dsn (key url) for project

```-bup=""``` // Optional, Type: string, betteruptime heartbeat url, eg. "https://betteruptime.com/api/v1/~~~"

```-message=""``` // Not needed at all, Type: string, optional commit message to pass for manual runs
