all: server

server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bodhi example/main.go