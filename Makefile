help: ## You are here! showing all command documenentation.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yaml

#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up: ## Run migration up
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate up

migrate-up-force: ## Run migrations UP force
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate force $(version)

migrate-down: ## Rollback latest one migration
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate down 1

migrate-down-all: ## Rollback all migration
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate down $(total)

migrate-create: ## Create migration
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --user=${USER} --rm migrate create -ext sql -seq -dir /migrations $(name)

#=======================#
#== SETUP ENVIRONMNET ==#
#=======================#

environment: ## Setup environment.
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d

shell-db: ## Enter to database console
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres

server: ## Run server.
	go run cmd/main.go

lint: ## Running golangci-lint for code analysis.
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm lint golangci-lint run -v