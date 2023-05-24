
# Mini Project: Financial Management API
Simple finance management application intended for bootcamp evaluation at [Studi Devsecops](https://studidevsecops.com/)

## Tech Stack

1. [Go](https://go.dev/)
2. [GIN Web Framework](https://github.com/gin-gonic/gin)
3. [PostgreSQL](https://www.postgresql.org/)
4. [Docker](https://www.docker.com/)
5. [Swagger](https://swagger.io/)

## Running Locally

1. `make environment`
2. `make migration-up`
3. Create `app.env` based on `app.env.sample`
4. Update `app.env`
5. `make server`
6. Application is running!

## Available Commands
```bash
help                           showing all command documenentation 

environment                    setup environment
lint                           run golangci-lint for code analysis
migrate-create                 create migration
migrate-down-all               rollback all migration
migrate-down                   rollback latest one migration
migrate-up                     run migration up
migrate-version                run migration to specific version
server                         run the server
shell-db                       enter to database console
```

## Example Users
Username:
    1. admin
    2. user_1
Password: satudua3