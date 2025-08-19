# Set up development environment
setup:
	@lefthook install
	@echo "✅ Development environment ready"

# Run Web BFF application
dev:
	@echo "🚀 Starting Web BFF..."
	@go run ./...

# Unit tests
test:
	@echo "🧪 Testing Web BFF..."
	@go test -v -race ./...
	@echo "✅ Tests passed"

# Format Go code using golangci-lint
fmt:
	@echo "🔧 Formatting Go code..."
	@golangci-lint fmt -c ./golangci-lint.yml
	@echo "✅ Code formatting complete"

# Run linter checks using gloangci-lint
lint:
	@echo "🔍 Running linter checks..."
	@golangci-lint run -c ./golangci-lint.yml
	@echo "✅ Linting complete"

# Run Swagger UI server
swagger:
	@echo "🚀 Starting Swagger UI server..."
	@./scripts/run-swagger-ui.sh

# Generate DTO types from OpenAPI components/schemas using oapi-codegen
# Note: Suppresses OpenAPI 3.1.x compatibility warning since we only use basic features
gen:
	@echo "🏗️  Generating DTO types from OpenAPI schema..."
	@oapi-codegen -config oapi-codegen.yml api/openapi.yml 2>&1 | (grep -v "You are using an OpenAPI 3.1.x specification" || true)
	@echo "✅ DTO generation complete"

# Validate OpenAPI specification
validate-api:
	@echo "🔍 Validating OpenAPI specification..."
	@npx @redocly/cli@2.0.6 lint api/openapi.yml
	@echo "✅ OpenAPI specification is valid"

# CI Build discarding artefacts
check-build:
	@echo ⏳ "Building..."
	@go build -o /dev/null ./...
	@echo "✅ Building complete"

.PHONY: setup dev test fmt lint swagger gen validate-api check-build