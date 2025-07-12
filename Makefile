APP_NAME := vyapaar
DB_FILE := vyapaar.db

# Run the app
run:
	go run ./cmd/server/main.go

# Build the app
build:
	go build -o $(APP_NAME) ./cmd/server

# Format the code
fmt:
	go fmt ./...

# Tidy dependencies
tidy:
	go mod tidy

# Test all packages
test:
	go test ./...

# Remove the SQLite DB (for reset/testing)
reset-db:
	rm -f $(DB_FILE)

# Recreate the database schema (assuming schema.sql exists)
migrate:
	sqlite3 $(DB_FILE) < migrations/schema.sql

# Clean binary and DB
clean:
	rm -f $(APP_NAME) $(DB_FILE)

# Help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run         Run the app"
	@echo "  build       Build the app binary"
	@echo "  fmt         Format Go code"
	@echo "  tidy        Clean and sync go.mod"
	@echo "  test        Run tests"
	@echo "  reset-db    Delete the SQLite database"
	@echo "  migrate     Apply DB schema from migrations/schema.sql"
	@echo "  clean       Remove binary and DB file"
