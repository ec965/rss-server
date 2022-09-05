package main

import (
	"github.com/ec965/rss-server/pkgs/models"
	"log"
)

func main() {
	db := models.Init("test.db")
	defer db.Close()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	models.Migrate()
}
