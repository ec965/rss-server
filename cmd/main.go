package main

import (
	"database/sql"

	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	version := ""
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	log.Print("connected to db")

	defer db.Close()

	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(version)

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
