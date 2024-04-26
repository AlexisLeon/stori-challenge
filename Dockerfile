# Dependencies Container
FROM golang:1.21 as build
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /go/src/app
RUN go build -o /go/bin/app main.go

FROM gcr.io/distroless/static
RUN adduser -D -u 1000 stori

COPY --from=build /go/bin/app /

USER stori
ENTRYPOINT ["/app"]
