TARGET=stori-api
.PHONY: build

build:
	GOOS=linux go build -o $(TARGET) ./main.go

test: build
	go test -v ./...
