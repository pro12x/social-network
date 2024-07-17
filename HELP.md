# Importantes Commandes
### Migrations
- Install: `go get -u github.com/golang-migrate/migrate/cmd/migrate`
- Install driver: `go install -tags '[DRIVER_NAME]' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
    - Exemple: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- Create migration: `migrate create -ext sql -dir [path_to_store_migrations_files] -seq [migration_name]`
    - Exemple: `migrate create -ext sql -dir ./migrations -seq create_users_table`
- Run migrations: `migrate -path [path_to_store_migrations_files] -database [connection_string] up`
    - Exemple: `migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up`