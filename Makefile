-include .env
export

# The default command to run when 'make' is called without arguments.
.DEFAULT_GOAL := help

# --- DEVELOPMENT COMMANDS (uses docker-compose.yml) ---

# Starts all services in FOREGROUND with watch mode for LOCAL development.
up:
	@echo "Starting all development services with watch mode enabled..."
	@echo "Press Ctrl+C to stop."
	@docker-compose up --build --watch

# Stops development containers but LEAVES data volumes intact.
down:
	@echo "Stopping development services and removing containers..."
	@docker-compose down

# --- DEPLOYMENT COMMANDS (uses deploy-compose.yml) ---

# Deploys the application to a server using the production-ready compose file.
deploy:
	@echo "Deploying application using deploy-compose.yml..."
	@docker-compose -f deploy-compose.yml up --build -d
	@echo "Deployment started in the background. Run 'make deploy-logs' to see output."

# Stops the production-deployed containers.
deploy-down:
	@echo "Stopping production containers defined in deploy-compose.yml..."
	@docker-compose -f deploy-compose.yml down

# Follows the logs of the production app container.
deploy-logs:
	@echo "Following logs for the deployed 'app' service..."
	@docker-compose -f deploy-compose.yml logs -f app

# --- SHARED & UTILITY COMMANDS ---

# Stops containers AND DELETES all data volumes. Use for a full reset.
nuke:
	@echo "WARNING: Stopping all services and DELETING ALL DATA..."
	@docker-compose down --volumes
	@docker-compose -f deploy-compose.yml down --volumes # Also nuke production volumes if present

# --- NEW: Applies migrations using our Go CLI ---
# This command runs our Go program with the '-migrate' flag.
# It requires that the database is running.
migrate-up:
	@echo "Applying database migrations using the Go CLI..."
	@go run ./cmd/cli -migrate

# --- NEW: Creates the initial admin user using our Go CLI ---
# This command runs our Go program with the '-seed-admin' flag.
# It's idempotent and safe to run multiple times.
init:
	@echo "Seeding initial admin user using the Go CLI..."
	@go run ./cmd/cli -seed-admin

# Build or rebuild service images using the DEV compose file
build:
	@echo "Building development images..."
	@docker-compose build

# Display help message
help:
	@echo ""
	@echo "Usage: make [command]"
	@echo ""
	@echo "--- Local Development ---"
	@echo "  up             - Start all services with live-reload for local dev."
	@echo "  down           - Stop local dev services."
	@echo ""
	@echo "--- Server Deployment ---"
	@echo "  deploy         - Deploy application to server using deploy-compose.yml."
	@echo "  deploy-down    - Stop the deployed application on the server."
	@echo "  deploy-logs    - Follow logs for the deployed application."
	@echo ""
	@echo "--- Database & Utilities ---"
	@echo "  migrate-up     - Apply database migrations using the Go CLI."
	@echo "  init           - Create the initial admin user using the Go CLI."
	@echo "  nuke           - Stop services AND DELETE ALL DATA (dev and prod)."
	@echo ""

# Declares which commands are "phony" (i.e., not actual files).
.PHONY: up down nuke deploy deploy-down deploy-logs migrate-up init build help
