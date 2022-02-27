# backup
![Yep](https://socialify.git.ci/joshhhhyyyy/Backup/image?descriptionEditable=backup%20a%20database%20using%20git.%20Err%20handling%20by%20Sentry%2C%20Reporting%20by%20Betteruptime.%20Made%20with%20%F0%9F%A9%B8%20%2C%20%F0%9F%98%93%20%26%20%F0%9F%98%AD&font=Source%20Code%20Pro&language=1&owner=1&pattern=Overlapping%20Hexagons&theme=Dark)

## What is this?
**Simple program to automatically backup a database using git. Err handling by Sentry, Reporting by Betteruptime. Made with 🩸 , 😓 &amp; 😭**

## Usage
```go get github.com/joshhhhyyyy/backup```

```go install github.com/joshhhhyyyy/backup```

```export PATH=$PATH:$(go env GOPATH)/bin``` (Add gopath to path)

```backup -key=""```

## Optional flags
**note: both single minus "-" and double minus "--" work fine

-bup="" // Type: string, betteruptime heartbeat url, eg. "https://betteruptime.com/api/v1/~~~"

-message="" // Type: string, optional commit message to pass for manual runs
