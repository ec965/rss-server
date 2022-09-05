package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ec965/rss-server/pkgs/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := models.Init("test.db")
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Print("Listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}
