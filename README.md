# teamOFP

## Components

* Remote - Spotify.app control app (run locally)
* API - Portal (manages queue, controls Remote)
* Amazon SQS - Messaging between Portal and Remote

## Getting started

### Remote

#### Build & Run

1. `go get` - Install required packages
1. `go build` - Build binary
1. `./remote` - Run remote app

*Note:* You will need to provide Remote with an Amazon access key and secret key for an SQS queue as environment variables or via a `.env` file.
```
ACCESS_KEY=XXXXXXXXXXX
ACCESS_SECRET=XXXXXXXXXXXXXXX
```

### API

#### Build & Run

1. `go get` - Install required packages
1. `go build` - Build binary
1. `./api` - Run Portal

*Note:* You will need to provide Remote with an Amazon access key and secret key for an SQS queue as environment variables or via a `.env` file.
```
ACCESS_KEY=XXXXXXXXXXX
ACCESS_SECRET=XXXXXXXXXXXXXXX
```

