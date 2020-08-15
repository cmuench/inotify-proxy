.PHONY: dev

format:
	go fmt ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o inotify-proxy inotify-proxy.go

test:
	GOOS=linux GOARCH=amd64 go test -v ./...

all: format test linux
