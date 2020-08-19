package main

import (
	"github.com/dfrnks/go-rest-api/internal"
	"github.com/dfrnks/go-rest-api/internal/Database"
	"log"
	"net/http"
)

func main() {
	Database.SyncDataBase()

	log.Fatal(http.ListenAndServe(":8000", internal.NewRouter()))
}
