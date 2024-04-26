# Dependencies Container
FROM golang:1.21-alpine as build
RUN apk add make
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /go/src/app

# pull deps
COPY go.mod go.sum ./
# RUN make deps

# build
COPY . /go/src/app/
RUN make build

FROM gcr.io/distroless/static:nonroot
COPY --from=build /go/src/app /
ENTRYPOINT ["/stori-api"]
