# Create dotenv file, on multiple lines
echo "PORT=1111" > .env
echo "DB_DRIVER=sqlite3" >> .env
echo "DB_CONNECTION=./pkg/db/mydb.db" >> .env
