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
	DatumZavodu string `json:"datum_zavodu"`
	Icon        string `json:"icon"`
}

//NejblizsiZavody vrací seznam 4 nejbližšéch závodů
func NejblizsiZavody(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	var nejblizsiZavody []NejblizsiZavod

	sql1 := "SELECT " + SqlZavody + ".id_zavodu," + SqlZavody + ".nazev_zavodu," + SqlZavody + ".kod_zavodu,DATE_FORMAT(" + SqlZavody + ".datum_zavodu,'%e.%c.%Y') AS datum_zavodu,typ_zavodu.icon FROM " + SqlZavody + ",typ_zavodu WHERE " + SqlZavody + ".datum_zavodu  > CURDATE() AND " + SqlZavody + ".typ_zavodu = typy_zavodu.id_typu_zavodu ORDER BY datum_zavodu ASC LIMIT 0,4"

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close() přijde to tu?
	for results.Next() {
		var nejblizsiZavod NejblizsiZavod
		err = results.Scan(&nejblizsiZavod.IDZavodu, &nejblizsiZavod.NazevZavodu, &nejblizsiZavod.KodZavodu, &nejblizsiZavod.DatumZavodu, &nejblizsiZavod.Icon)
		nejblizsiZavody = append(nejblizsiZavody, nejblizsiZavod)
	}
	//Headers(w)
	json.NewEncoder(w).Encode(nejblizsiZavody)
}
