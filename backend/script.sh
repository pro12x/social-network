# Create dotenv file, on multiple lines
echo "PORT=1111" > .env
# shellcheck disable=SC2129
echo "DB_DRIVER=sqlite3" >> .env
echo "DB_CONNECTION=./pkg/db/mydb.db" >> .env
echo "DB_MIGRATION_PATH=./pkg/db/migrations/sqlite" >> .env
echo "DEFAULT_API_LINK=/api/v1/social-network" >> .env
echo "METHOD_NOT_ALLOWED=\"Method not allowed\"" >> .env
echo "NOT_FOUND=\"Not found\"" >> .env
echo "INVALID_CREDS=\"Invalid credentials\"" >> .env
echo "CONTENT_TYPE=\"Content-Type\"" >> .env
echo "APPLICATION_JSON=\"application/json\"" >> .env
echo "USER_ID_REQUIRED=\"User id is required\"" >> .env
