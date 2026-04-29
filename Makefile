include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker-compose up -d todoapp-postgres
env-dow:
	docker-compose down todoapp-postgres
env-cleanup:
	read -p "Are you sure you want to cleanup the environment? (y/n): " confirm; \
	if [ "$$confirm" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf out/pgdata && \
		echo "Environment cleaned up successfully"; \
	else \
		echo "Environment cleanup cancelled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Migration name is required"; \
		exit 1; \
	fi; \
	MSYS_NO_PATHCONV=1 docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-actions action=up
migrate-down:
	@make migrate-actions action=down
migrate-actions:
	if [ -z "$(action)" ]; then \
		echo "Action is required"; \
		exit 1; \
	fi; \
	MSYS_NO_PATHCONV=1 docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		"$(action)"

env-socat:
	docker compose up -d todoapp-env-alpine-socat