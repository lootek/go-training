FROM golang

WORKDIR /go/src/go-training

# COPY . .
#RUN go get -d -v ./...
#RUN go install -v ./...

RUN apt-get update && apt-get install --yes --auto-remove \
	apt-utils \
	curl \
	file \
	git \
	nano \
	net-tools \
	tree \
	&& rm -rf /var/lib/apt/lists/*

RUN curl https://glide.sh/get | sh
RUN go get -v -u golang.org/x/tools/cmd/goimports
RUN go get -v -u github.com/derekparker/delve/cmd/dlv
RUN go get -v -u github.com/cespare/reflex

ENTRYPOINT bash
CMD ["tail", "-f", "/dev/null"]
