#!/bin/bash

docker run \
	--interactive \
	--tty \
	--rm \
	--publish 8000:80 \
	--publish 8008:8333 \
	--security-opt=seccomp:unconfined \
	--cap-add=SYS_PTRACE \
	--volume "$PWD":/go/src/go-training \
	--workdir /go/src/go-training \
	go-training
