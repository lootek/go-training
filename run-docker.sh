#!/bin/bash

docker run -it --rm -v "$PWD":/go/src/go-training -w /go/src/go-training go-training

