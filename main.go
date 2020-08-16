package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db, err = sql.Open("sqlite3", "./database.sqlite")

func main() {
	if err != nil {
		panic(err)
	}

	syncDataBase()

	log.Fatal(http.ListenAndServe(":8000", NewRouter()))
}
