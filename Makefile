.DEFAULT_GOAL := help
.PHONY: help
help: ## Display this help
	@grep -h "##" $(MAKEFILE_LIST) | grep -v grep | sed -e 's/\\$$//' | sed -e 's/##//' | column -t -s ':'


# Variables
LOCAL_BIN := $(CURDIR)/bin
PROTO_DIR = proto
GEN_DIR = gen
BUF = buf
PROTOC_GEN_GO = $(LOCAL_BIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(LOCAL_BIN)/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY = $(LOCAL_BIN)/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV2 = $(LOCAL_BIN)/protoc-gen-openapiv2
PROTOC_GEN_GO_VTPROTO = $(LOCAL_BIN)/protoc-gen-go-vtproto
PROTOC_GEN_VALIDATE = $(LOCAL_BIN)/protoc-gen-validate
PROTOC_GEN_GOCLAY = $(LOCAL_BIN)/protoc-gen-goclay

# Targets
.PHONY: all clean generate install-plugins

all: generate

clean:
	rm -rf $(GEN_DIR)

generate: install-buf ## Generate code
	$(BUF) generate --template buf.gen.yaml
	

# Install buf
install-buf:
	if ! command -v buf &> /dev/null; then \
		brew install bufbuild/buf/buf; \
	fi

# Install protoc-gen-go
$(PROTOC_GEN_GO):
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install protoc-gen-go-grpc
$(PROTOC_GEN_GO_GRPC):
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install protoc-gen-grpc-gateway
$(PROTOC_GEN_GRPC_GATEWAY):
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

# Install protoc-gen-openapiv2
$(PROTOC_GEN_OPENAPIV2):
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# Install protoc-gen-go-vtproto
$(PROTOC_GEN_GO_VTPROTO):
	GOBIN=$(LOCAL_BIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

# Install protoc-gen-validate
$(PROTOC_GEN_VALIDATE):
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest

$(PROTOC_GEN_GOCLAY):
	GOBIN=$(LOCAL_BIN) go install github.com/utrack/clay/cmd/protoc-gen-goclay@latest

# Ensure all plugins are installed
install-plugins: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY) $(PROTOC_GEN_OPENAPIV2) $(PROTOC_GEN_GO_VTPROTO) $(PROTOC_GEN_VALIDATE) $(PROTOC_GEN_GOCLAY) ## Install all required protoc plugins
	@echo "All plugins installed successfully"
	$(MAKE) install-lint


.PHONY: migrate-local-up migrate-local-down

# Database connection settings
DB_USER := user
DB_PASSWORD := password
DB_HOST := localhost
DB_PORT := 5432
DB_NAME := mydatabase
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate-local-up: ## Run database migrations up
	goose -dir scripts/migrations/postgres postgres "$(DB_URL)" up

migrate-local-down: ## Roll back database migrations
	goose -dir scripts/migrations/postgres postgres "$(DB_URL)" down


.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run ./...

install-lint: ## Install golangci-lint
	brew install golangci-lint