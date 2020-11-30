package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//NejblizsiZavod je struktura pro func NejblizsiZavody
type NejblizsiZavod struct {
	IDZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
	KodZavodu   string `json:"kod_zavodu"`
}

//NejblizsiZavody vrací seznam 4 nejbližšéch závodů
func NejblizsiZavody(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	var nejblizsiZavody []NejblizsiZavod

	sql1 := "SELECT id_zavodu,nazev_zavodu,kod_zavodu FROM " + SqlZavody + " WHERE datum_zavodu  > CURDATE() ORDER BY datum_zavodu ASC LIMIT 0,4"

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close() přijde to tu?
	for results.Next() {
		var nejblizsiZavod NejblizsiZavod
		err = results.Scan(&nejblizsiZavod.IDZavodu, &nejblizsiZavod.NazevZavodu, &nejblizsiZavod.KodZavodu)
		nejblizsiZavody = append(nejblizsiZavody, nejblizsiZavod)
	}
	//Headers(w)
	json.NewEncoder(w).Encode(nejblizsiZavody)
}
