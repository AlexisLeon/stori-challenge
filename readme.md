# Technical challenge for storicard.com

## Getting started

These instructions assume you're running mac os.

Run `scripts/bootstrap.sh` or follow the following instructions:

### Install the required packages
- go
- postgres
- golang-migrate
- golangci-lint
- mockery

Install dependencies:
```shell
$ brew install go postgres golang-migrate golangci-lint mockery
```

Run Postgres as background service: `brew services start postgresql`
Or if you don't want/need a background service you can just run: `pg_ctl -D /usr/local/var/postgres start`

### Run the project

You can manually compile the project with `make build` and run the binary `./stori-api`

## Testing

```
$ go test
```
