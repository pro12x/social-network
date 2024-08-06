# Variables
BACKEND_DIR = backend
FRONTEND_DIR = frontend
BACKEND_APP_NAME = backend
DEPENDENCIES_FILE = main.sh

# Targets
build-backend:
	@echo "Building backend"
	@cd $(BACKEND_DIR) && go build -o $(BACKEND_APP_NAME) server.go

run-dependencies:
	@echo "Installing dependencies"
	@cd $(BACKEND_DIR) && chmod +x $(DEPENDENCIES_FILE) && ./$(DEPENDENCIES_FILE)

run-backend: $(BACKEND_APP_NAME)
	@echo "Running backend"
	@cd $(BACKEND_DIR) && ./$(BACKEND_APP_NAME)

run-frontend:
	@echo "Running frontend"
	@cd $(FRONTEND_DIR) && npm install
	@cd $(FRONTEND_DIR) && nohup ng serve --port 1112

run-all: run-dependencies build-backend run-frontend run-backend
	@echo "Running all"

clear:
	@echo "Clearing"
	@cd $(BACKEND_DIR) && rm -rf $(BACKEND_APP_NAME)
	@cd $(FRONTEND_DIR) && rm -rf node_modules
	@cd $(FRONTEND_DIR) && rm -rf dist