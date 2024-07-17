# Create dotenv file, on multiple lines
echo "PORT=1111" > .env
# shellcheck disable=SC2129
echo "DB_DRIVER=sqlite3" >> .env
echo "DB_CONNECTION=./pkg/db/mydb.db" >> .env
echo "DB_MIGRATION_PATH=./pkg/db/migrations/sqlite" >> .env
echo "DEFAULT_API_LINK=/api/v1/social-network" >> .env
