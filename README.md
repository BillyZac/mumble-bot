# Mumblee Little Bot
Server that takes post data and forwards it to Twitter to be posted as a user's tweet.

## Build and install
Make sure you've set your GOPATH:
```
$ export GOPATH=$HOME/work
```
 
Then build and install using the `go` tool:
```
$ go install github.com/BillyZac/mumble-bot
```

To start the bot:
```
$ mumble-bot
```
This runs a web server on PORT 8080, which listens for GET requests and then "responds" by firing off a Tweet.

Note, if the executable doesn't run, double-check that your workspace's `bin` is on your PATH:
```
$ echo PATH
```

If it's not, add it:
```
$ export GOPATH/$HOME/work
```

For more on getting started with Go, [see the Go docs](https://golang.org/doc/code.html).

## Run tests
This command runs any file whose name follows the pattern *_test.go. See https://golang.org/cmd/go/#hdr-Test_packages.

```
$ go test
```

## Contributing
Dependencies are vendored in the vendor directory
`go get -u github.com/kardianos/govendor` to work with dependencies

