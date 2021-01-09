.PHONY: dev

LDFLAGS="-ldflags=-s -w -X 'main.Version=dev-linux'"

format:
	go fmt ./...

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o inotify-proxy-linux-amd64 $(LDFLAGS) inotify-proxy.go

build-linux-arm64:
	GOOS=linux GOARCH=amd64 go build -o inotify-proxy-linux-arm64 $(LDFLAGS) inotify-proxy.go

test:
	go test -v ./...

build: build-linux-amd64 build-linux-arm64
all: format test build
