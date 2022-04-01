deps:
	go mod tidy

build:
	go build -o bin/quik-gaming-interview cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...