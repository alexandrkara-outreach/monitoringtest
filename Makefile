all: build


build:
	go build ./...

run:
	go run cmd/example-monitoring/main.go
