TARGET=stori
.PHONY: build deps migrate

build:
	go build -o $(TARGET) ./main.go

deps:
	@go mod download
	@go mod verify

test: build
	go test -v ./...

docker-build:
	docker build . -t $(TARGET)
	docker tag $(TARGET) $(TARGET):latest
