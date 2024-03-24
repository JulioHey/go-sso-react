all: build

build:
	go build -o build/authorizer cmd/main.go

run:
	go run cmd/main.go

test:
	go test -count=1 ./...