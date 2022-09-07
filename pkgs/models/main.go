package models

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(dataSourceName string) *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate(sourceUrl string) {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(sourceUrl, "sqlite3", instance)

	if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
		log.Println(err)
	} else if err != nil {
		log.Fatalln(err)
	}
}
