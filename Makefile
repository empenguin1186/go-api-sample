.PHONY: run codegen docker-compose create-migrations run-migration

include .env

run:
	go run cmd/server.go

codegen:
	docker run --rm -v "$(CURDIR):/app" openapitools/openapi-generator-cli generate \
	-i "/app/spec/openapi.yaml" \
	-g go-server \
	-o /app/code_generated

docker-compose:
	docker-compose up -d --remove-orphans

create-migrations:
	docker run -v "$(CURDIR)/migrations:/migrations" migrate/migrate \
	create -ext sql -dir migrations -seq create_favorite_table

OPERATION = up
MIGRATION_NUM = 1
run-migration:
	docker run -v "$(CURDIR)/migrations:/migrations" --network host migrate/migrate \
	-path=/migrations/ -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:15432/sample?sslmode=disable" $(OPERATION) $(MIGRATION_NUM)