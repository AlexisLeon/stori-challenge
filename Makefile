TARGET=stori-api
.PHONY: build deps

build:
	GOOS=linux go build -o $(TARGET) ./main.go

deps:
	@go mod download
	@go mod verify

test: build
	go test -v ./...

docker-build:
	docker build .
