# Importantes Commandes
### Migrations
- Install: `go get -u github.com/golang-migrate/migrate/cmd/migrate`
- Install driver: `go install -tags '[DRIVER_NAME]' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
    - Exemple: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- Create migration: `migrate create -ext sql -dir [path_to_store_migrations_files] -seq [migration_name]`
    - Exemple: `migrate create -ext sql -dir ./migrations -seq create_users_table`
- Run migrations: `migrate -path [path_to_store_migrations_files] -database [connection_string] up`
    - Exemple: `migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up`

### Project
- The script.sh script is used to set the environment variables.
- The main.sh script is used to initialize the project module and import the necessary packages.
- The Dockerfile is used to build the Docker image.
- The docker-compose.yml file is used to run the Docker container.
- The .gitignore file is used to ignore the files and directories that are not required in the Git repository.
- The .env file is used to store the environment variables.
- The backend directory contains the Go code for the backend service.
- The frontend directory contains the React code for the frontend service.
- The Makefile is used to automate the build process.
- The README.md file contains the instructions for building and running the project.
- The backend/server.go file contains the main Go code for the backend service.
- The backend/go.mod file contains the Go module definition for the backend service.