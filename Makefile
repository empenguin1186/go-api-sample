.PHONY: run codegen

run:
	go run cmd/server.go

codegen:
	docker run --rm -v "$(CURDIR):/app" openapitools/openapi-generator-cli generate \
	-i "/app/spec/openapi.yaml" \
	-g go-server \
	-o /app/code_generated
