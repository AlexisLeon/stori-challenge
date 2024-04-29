# Technical challenge for storicard.com

- The project was built as a common line interface. You can execute the following commands.
  - `migrate`
  - `settlement`
- The settlement process assumes that we only have one user and only one account
- Each transaction is saved to the db and the account balance gets updated
- Normally, we would create a ledger to keep track of all the transactions and the balances. We're not doing that for simplicity

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

## Docker

The docker container uses an ENTRYPOINT to load the binary and can take commands (`CMD`) as arguments.
This allows us to run the container with `docker run stori <command>`

## Configure the project

### Database config

Configure your database credentials in `config.yml`.

### Mailer config

The mailer is configured to use AWS SES. You'll need to set the following environment variables:

When running docker, make sure to pass the env variables as args (`ARG`).

| Variable             |
|----------------------|
| `AWS_REGION`           |
| `AWS_ACCESS_KEY_ID`    |
| `AWS_SECRET_ACCESS_KEY` |

- The credentials need to have access to the `ses:*` action.
- Update the `config.yml` file with the email you want to send from.
- Make sure that the email you're using is verified in AWS SES.

> To change the destination email, update the `email` field in `db/migrations/000004_create_test_user.up.sql`

```diff
INSERT INTO {{ index .Options "Namespace" }}.users VALUES (
'00000000-0000-0000-0000-000000000000',
+ 'myemail@storicard.com'
);
```

## Compile the project

You can manually compile the project with `make build` and run the binary `./stori`
You can also run the docker container with `make docker-build`. This will build docker image and tag it as `stori:latest`.

## Run migrations

Run the migrations with `stori migrate` or `docker run stori migrate`

## Run the project

Run the project with `stori settlement`. This will process the `settlement.csv` file send an email with the settlement report to the user.

> Note that you'll need to mount the `settlement.csv` file to the container as is not included in the image.
