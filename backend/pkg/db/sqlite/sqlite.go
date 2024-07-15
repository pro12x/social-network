package sqlite

import (
	"backend/pkg"
	"database/sql"
	"log"
	"os"
)

type Database struct {
	db *sql.DB
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) Close() {
	err := d.GetDB().Close()
	if err != nil {
		log.Printf("Error closing database connection\nCaused by: %v", err)
		return
	}
}

func Connect() (*Database, error) {
	err1 := pkg.Environment()
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONNECTION"))
	if err != nil || err1 != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection\nCaused by: %v", err)
		}
		return nil, err
	}
	log.Println("Connected to the database")
	return &Database{db: db}, nil
}
