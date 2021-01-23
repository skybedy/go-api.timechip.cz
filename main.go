package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go-api.timechip.cz/conf"
	"go-api.timechip.cz/db"
	"go-api.timechip.cz/routes"
)

//var db *sql.DB
var err error

const Port = "1313"

func main() {

	db.Mdb, err = sql.Open(conf.DbDriver, conf.DbUser+":"+conf.DbPass+"@/"+conf.DbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Mdb.Close()

	router := routes.NewRouter()

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
