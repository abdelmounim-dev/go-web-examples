package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "you requested book %s, page %s", vars["title"], vars["page"])
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn't load env vars")
	}

	connStr := os.Getenv("CONN")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Couldn't connect to the db", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db connection not working", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS users (
      id serial primary key,
      username text not null,
      password text not null,
      created_at date
    );
  `

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("error creating users table: ", err)
	}

	http.ListenAndServe(":3000", router)
}
