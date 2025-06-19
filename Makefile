.PHONY: generate-swagger

SWAGGER_HOST ?= localhost:8080

generate-swagger:
	sed -i "s|@host .*|@host $(SWAGGER_HOST)|" cmd/main.go