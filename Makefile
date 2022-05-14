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
.SILENT:
run-migration:
	docker run -v "$(CURDIR)/migrations:/migrations" --network host migrate/migrate \
	-path=/migrations/ -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" $(OPERATION) $(MIGRATION_NUM)

FILE_DATA_NAME = sample.data.dump
database-data-dump:
	docker run -v "$(CURDIR)/dump:/dump" -it --rm --network host postgres:${POSTGRES_VERSION} pg_dump -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -v -a -f /dump/$(FILE_DATA_NAME) -W

FILE_SCHEMA_NAME = sample.schema.dump
database-schema-dump:
	docker run -v "$(CURDIR)/dump:/dump" -it --rm --network host postgres:${POSTGRES_VERSION} pg_dump -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -v -s -f /dump/$(FILE_SCHEMA_NAME) -W