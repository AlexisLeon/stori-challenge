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
WORKDIR /
COPY --from=build /go/src/app /
COPY config.yml /
COPY db/migrations /db/migrations

ARG AWS_ACCESS_KEY_ID
ENV AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
ARG AWS_REGION
ENV AWS_REGION=${AWS_REGION}
ARG AWS_SECRET_ACCESS_KEY
ENV AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}

ENTRYPOINT ["/stori"]
