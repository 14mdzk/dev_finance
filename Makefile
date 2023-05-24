help: ## showing all command documenentation
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yaml

#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up: ## run migration up
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate up

migrate-version: ## run migration to specific version
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate force $(version)

migrate-down: ## rollback latest one migration
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate down 1

migrate-down-all: ## rollback all migration
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate down $(total)

migrate-create: ## create migration
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate create -ext sql -seq -dir /migrations $(name)

#=======================#
#== SETUP ENVIRONMNET ==#
#=======================#

environment: ## setup environment
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d

shell-db: ## enter to database console
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres

server: ## run the server
	go run cmd/*.go

lint: ## run golangci-lint for code analysis
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm lint golangci-lint run -v