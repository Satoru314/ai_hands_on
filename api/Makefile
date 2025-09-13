.PHONY: run-dev init-db stop clean help

# Development server (Docker + Go)
run-dev:
	@echo "🚀 Starting development environment..."
	@docker-compose -f ../docker-compose.yml up -d
	@echo "✅ Development environment started!"
	@echo "📍 API server: http://localhost:8080"
	@echo "📍 Database: localhost:5432"

# Initialize database tables
init-db:
	@echo "🗄️  Initializing database..."
	@go run cmd/initdb/main.go
	@echo "✅ Database initialization completed!"

# Stop development environment
stop:
	@echo "🛑 Stopping development environment..."
	@docker-compose -f ../docker-compose.yml down

# Clean up (remove containers and volumes)
clean:
	@echo "🧹 Cleaning up..."
	@docker-compose -f ../docker-compose.yml down -v --remove-orphans

# Show help
help:
	@echo "Available commands:"
	@echo "  make run-dev  - Start Docker + Go development server"
	@echo "  make init-db  - Initialize database tables"
	@echo "  make stop     - Stop development environment"
	@echo "  make clean    - Stop and remove all containers/volumes"
	@echo "  make help     - Show this help message"