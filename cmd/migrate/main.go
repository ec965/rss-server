package main

import (
	"log"

	"github.com/ec965/rss-server/pkgs/env"
	"github.com/ec965/rss-server/pkgs/models"
)

func main() {
	dbUrl := env.Get("DATABASE_URL", "test.db")
	migrationsDir := env.Get("MIGRATIONS_DIR", "file://db/migrations")
	db := models.Init(dbUrl)
	defer db.Close()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	models.Migrate(migrationsDir)
}
