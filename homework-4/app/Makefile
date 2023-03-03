FLAGS?=-v

.PHONY: gen
gen: ## Generate projects files and components.
	@oapi-codegen -config oapi/server.cfg.yaml oapi/openapi.yaml
	@oapi-codegen -config oapi/models.cfg.yaml oapi/openapi.yaml
	@oapi-codegen -config oapi/spec.cfg.yaml oapi/openapi.yaml
	@oapi-codegen -config oapi/client.cfg.yaml oapi/openapi.yaml

.PHONY: build
build:
	docker build --tag chelout/otus-homework_4 .

.PHONY: push
push:
	docker push chelout/otus-homework_4

.PHONY: run
run:
	docker run -p 8000:8000 chelout/otus-homework_4
