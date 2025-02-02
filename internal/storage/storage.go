package storage

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"log"
	"os"
)

type Storage struct {
	db   *sql.DB
	Goqu *goqu.Database
}

func NewStorage(db *sql.DB) *Storage {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := "user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Can't connect to db: ", err)
	}

	return &Storage{
		db:   db,
		Goqu: goqu.New("postgres", db),
	}
}
