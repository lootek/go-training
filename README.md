# go-training

## Resources:

### Highly recommended reading:

#### Basics
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
`reflex -r '(\.go)$' -s -- bash -c 'go install go-training/server && $GOPATH/bin/server'`

### Building in docker
* https://hub.docker.com/_/golang
* https://blog.hasura.io/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c
