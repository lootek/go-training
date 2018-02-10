#!/bin/bash

docker run \
	--interactive \
	--tty \
	--rm \
	--publish 8000:80 \
	--publish 8008:8333 \
	--volume "$PWD":/go/src/go-training \
	--workdir /go/src/go-training \
	go-training
