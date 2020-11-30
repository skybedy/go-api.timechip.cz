package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

const Port = "1313"
const KodZavodu = "cc_mcr_a_pohar_cams_sikluv_mlyn_1"
const RaceYear = "2020"
const SqlZavody = "zavody_" + RaceYear
const SqlPodzavody = "podzavody_" + RaceYear
const SqlKategorie = "kategorie_" + RaceYear
const SqlOsoby = "osoby"

var SqlZavod = "zavod_" + KodZavodu + "_" + RaceYear

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fungujem")
	fmt.Printf("%v\n", "tady")
	//json.NewEncoder(w).Encode("Still alive!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "timechip API")
		//fmt.Fprintf(os.Stdout, "timechip API stdout")
		fmt.Println(r)

	}).Methods("GET")
	router.HandleFunc("/homepage/nejblizsi-zavody", NejblizsiZavody).Methods("GET")

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
