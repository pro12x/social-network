package sqlite

import (
	"backend/pkg/utils"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Database struct
type Database struct {
	db *sql.DB
}

// GetDB function
func (d *Database) GetDB() *sql.DB {
	return d.db
}

// Close function
func (d *Database) Close() {
	err := d.GetDB().Close()
	if err != nil {
		log.Printf("Error closing database connection\nCaused by: %v", err)
		utils.LoggerError.Println(utils.Error+"Error closing database connection caused by: %v %v", err.Error(), utils.Reset)
		return
	}
}

// Connect function
func Connect() (*Database, error) {
	err1 := utils.Environment()
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONNECTION"))
	if err != nil || err1 != nil {
		return nil, errors.New("error connecting to the database: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection\nCaused by: %v", err)
			utils.LoggerError.Println(utils.Error+"Error closing database connection caused by: %v %v", err.Error(), utils.Reset)
		}
		return nil, err
	}
	log.Println("Connected to the database")
	utils.LoggerInfo.Println(utils.Info + "Connected to the database " + utils.Reset)
	return &Database{db: db}, nil
}

// Migrate function
func Migrate(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return errors.New("error creating sqlite3 driver: " + err.Error())
	}

	//
	m, err := migrate.NewWithDatabaseInstance("file://"+os.Getenv("DB_MIGRATION_PATH"), "sqlite3", driver)
	if err != nil {
		return errors.New("error creating migration instance: " + err.Error())
	}

	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.New("error rolling back migration: " + err.Error())
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.New("error applying migration: " + err.Error())
	}

	log.Println("Database migrated")
	utils.LoggerInfo.Println("Database migrated")
	return nil
}
