# go-training

## Purpose
This is not intended to be a complete Golang overview but rather an intro 
for developers already using other languages. 

It's kind of a background for my short talk + live coding session 
in which I wanted to focus on some of the very Golang-specific aspects.

## Idea
The project consists of 2 parts:
* an HTTP server (with embedded in-memory storage for int-s)
* an HTTP client (for communication with the server)

Both are capable of handling both XML and JSON in a basic way.

## Disclaimer
The code has some drawbacks from the architectural point of view. 
It also doesn't make sense to have a server storing int-s in that way.  

## Quick start

1. Clone this repo, `cd` to it
2. `$ docker build -t go-training .`
3. `$ ./run-docker.sh`
4. Inside the docker run any of:
   * `> ./watch.sh server` for live-reload experience of the server part
   * `> ./watch.sh cli` for live-reload experience of the CLI part
   * `> go install && $GOPATH/bin/<binary>` to run any of the compiled/installed binaries

   or basically do anything with the `go` tool.

## Golang Resources:

### Highly recommended reading:

#### Basics
* https://play.golang.org
* https://tour.golang.org
* https://golang.org/cmd/go
* https://golang.org/doc/faq

#### A must-read
* https://golang.org/doc/effective_go.html

#### Do as Gophers do
* https://github.com/golang/go/wiki/CodeReviewComments
* https://pocketgophers.com/idiomatic-go
* https://go-proverbs.github.io
* https://talks.golang.org/2014/readability.slide
* https://github.com/golang

#### How-to basics
* https://gobyexample.com
* http://tmrts.com/go-patterns

#### Docs
* https://golang.org/pkg
* https://godoc.org

#### Tips
* https://rakyll.org/go-tool-flags

#### Keep up-to-date
* https://blog.golang.org
* https://golangweekly.com

### List of Go frameworks, libraries and software
* https://awesome-go.com

### Nice tools
* Dependency management \
`curl https://glide.sh/get | sh` \
`go get -u github.com/golang/dep/cmd/dep`
* A `go fmt` replacement \
`$ go get -u golang.org/x/tools/cmd/goimports`
* Debugger \
`$ go get -u github.com/derekparker/delve/cmd/dlv`
* Live compilation \
`$ go get -u github.com/cespare/reflex` \
and then eg. \
`reflex -r '(\.go)$' -s -- bash -c 'go install <pkg> && $GOPATH/bin/<bin>'`

### Building in docker
* https://hub.docker.com/_/golang
* https://blog.hasura.io/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c
