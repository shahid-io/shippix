# === Development dependencies installer ===
# Installs or updates CLI tools used for linting, codegen, hot-reload, etc.
.PHONY: deps

deps:
	@echo "Installing dev tools…"
	# Linter
	@if ! command -v golangci-lint >/dev/null; then \
	  echo "→ Installing golangci-lint"; \
	  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.59.0; \
	else \
	  echo "golangci-lint already installed"; \
	fi
	# Proto codegen (buf)
	@if ! command -v buf >/dev/null; then \
	  echo "→ Installing buf"; \
	  curl -sSL https://github.com/bufbuild/buf/releases/download/v1.24.0/buf-Linux-x86_64 -o /usr/local/bin/buf && chmod +x /usr/local/bin/buf; \
	else \
	  echo "buf already installed"; \
	fi
	# Swagger generator (swag)
	@if ! command -v swag >/dev/null; then \
	  echo "→ Installing swag"; \
	  go install github.com/swaggo/swag/cmd/swag@latest; \
	else \
	  echo "swag already installed"; \
	fi
	# Hot-reload (air)
	@if ! command -v air >/dev/null; then \
	  echo "→ Installing air"; \
	  go install github.com/cosmtrek/air@latest; \
	else \
	  echo "air already installed"; \
	fi
	@echo "All dev tools are installed or up-to-date."

proto-gen:
	@echo "Generating gRPC and protobuf code..."
	# If using buf:
	buf generate --template buf.gen.yaml

	# OR, if using protoc directly, for each .proto:
	# protoc \
	#   --go_out=pkg --go_opt=paths=source_relative \
	#   --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
	#   -I proto proto/*.proto