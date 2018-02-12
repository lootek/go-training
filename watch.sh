#!/usr/bin/env bash

app=${1:-server}

reflex -r '(\.go)$' -s -- bash -c "go install github.com/lootek/go-training/$app && $GOPATH/bin/$app > /dev/stderr"
