package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go-api.timechip.cz/routes"
)

var db *sql.DB
var err error

const Port = "1313"

// KodZavodu
/*
var KodZavodu = "cc_mcr_a_pohar_cams_sikluv_mlyn_1"
var currentTime = time.Now()

var RaceYear = currentTime.Year()
var SqlPodzavody = "podzavody_" + strconv.Itoa(RaceYear)
var SqlKategorie = "kategorie_" + strconv.Itoa(RaceYear)
var SqlOsoby = "osoby"
var SqlZavody = "zavody_" + strconv.Itoa(RaceYear)
var SqlZavod = "zavod_" + KodZavodu + "_" + strconv.Itoa(RaceYear)
*/

func TestFunc1(RaceYear string, RaceCode int) {
	//fmt.Println(TestConst)
	fmt.Println(RaceYear + "," + strconv.Itoa(RaceCode))
}

var SqlZavodySc string
var SqlZavodyBc string

func SqlZkatky(RaceYear string) {
	SqlZavodySc = "zavody_" + RaceYear + " zd,"
	SqlZavodyBc = "zavody_" + RaceYear + " zd"
}

func main() {
	router := routes.NewRouter()

	db, err = sql.Open("mysql", "skybedy:mk1313life@tcp(127.0.0.1:3306)/timechip_cz?multiStatements=true")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
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
		//log.Fatal("main: couldn't start simple server: %v\n", err)
		//log.Fatal().Err(err)
	}

}
