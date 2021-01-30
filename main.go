package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go-api.timechip.cz/conf"
	"go-api.timechip.cz/db"
	"go-api.timechip.cz/routes"
)

var err error

const Port = "1312"

func main() {

	db.Mdb, err = sql.Open(conf.DbDriver, conf.DbUser+":"+conf.DbPass+"@/"+conf.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Mdb.Close()

	logFile, err := os.OpenFile(conf.AppPath+"/log/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	router := routes.NewRouter()

	port, ok := os.LookupEnv("PORT")

	if err != nil {
		panic(err.Error())
	}

	if !ok {
		port = Port
	}

	server := &http.Server{
		Handler: router,
		//Addr:    "127.0.0.1:" + port,
		Addr: "0.0.0.0:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Spouští se http server na portu", port)
	fmt.Println("Spouští se server na portu", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
