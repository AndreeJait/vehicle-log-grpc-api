include .env
export

CURRENT_DIR=$(shell pwd)


.PHONY: nQsd
nQsd:

.PHONY: gRpc
gRpc: gRpc-run
gRpc-run:
	@go run main.go gRpc

.PHONY: migrate
migrate: migrate-new migrate-down migrate-up migrate-fresh

migrate-new:
	@read -p "migration name: " MIGRATE_NAME; \
	CGO_ENABLED="0" go install github.com/rubenv/sql-migrate/...@latest; \
	sql-migrate new $${MIGRATE_NAME}
migrate-up:
	@go run main.go migrate up
migrate-down:
	@go run main.go migrate down
migrate-fresh:
	@go run main.go migrate fresh
