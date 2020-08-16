.PHONY: dev

LDFLAGS="-ldflags=-X 'main.Version=dev-linux'"

format:
	go fmt ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o inotify-proxy $(LDFLAGS) inotify-proxy.go

test:
	GOOS=linux GOARCH=amd64 go test -v ./...

all: format test linux
