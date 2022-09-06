package main

import (
	"log"
	"net/http"

	"github.com/ec965/rss-server/pkgs/env"
	"github.com/ec965/rss-server/pkgs/handlers"
	"github.com/ec965/rss-server/pkgs/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := env.Get("PORT", "3000")
	dbUrl := env.Get("DATABASE_URL", "test.db")

	db := models.Init(dbUrl)
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/feeds", handlers.GetFeeds)
	r.Get("/feeds/update", handlers.UpdateFeeds)
	r.Post("/feed", handlers.AddFeed)
	r.Delete("/feed", handlers.DeleteFeed)

	log.Print("Listening on http://localhost:" + port)

	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatal(err)
	}
}
