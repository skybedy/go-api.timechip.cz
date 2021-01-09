package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"go-api.timechip.cz/routes"
)

var db *sql.DB
var err error

const Port = "1313"

func main() {

	db, err = sql.Open("mysql", "skybedy:mk1313life@tcp(127.0.0.1:3306)/timechip_cz?multiStatements=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := routes.NewRouter(db)

	log.Flags()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = Port
	}

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("main: running simple server on port", port)
	if err := server.ListenAndServe(); err != nil {
	}
}
