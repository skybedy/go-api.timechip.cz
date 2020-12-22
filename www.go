package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//NejblizsiZavod je struktura pro func NejblizsiZavody
type NejblizsiZavod struct {
	IDZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
	KodZavodu   string `json:"kod_zavodu"`
	DatumZavodu string `json:"datum_zavodu"`
	Web         string `json:"web"`
	Icon        string `json:"icon"`
}

//NejblizsiZavody vrací seznam 4 nejbližšéch závodů
func NejblizsiZavody(w http.ResponseWriter, r *http.Request) {
	var nejblizsiZavody []NejblizsiZavod

	sql1 :=
		"SELECT z.id_zavodu,z.nazev_zavodu,z.kod_zavodu,DATE_FORMAT(z.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,z.web" +
			" FROM " + SqlZavody + " z,typ_zavodu tz" +
			" WHERE" +
			" z.datum_zavodu  > CURDATE() AND" +
			" z.typ_zavodu = tz.id_typ_zavodu" +
			" ORDER BY z.datum_zavodu ASC LIMIT 0,4"

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close() přijde to tu?
	for results.Next() {
		var nejblizsiZavod NejblizsiZavod
		err = results.Scan(&nejblizsiZavod.IDZavodu, &nejblizsiZavod.NazevZavodu, &nejblizsiZavod.KodZavodu, &nejblizsiZavod.DatumZavodu, &nejblizsiZavod.Icon, &nejblizsiZavod.Web)
		nejblizsiZavody = append(nejblizsiZavody, nejblizsiZavod)
	}
	pocetZavodu := len(nejblizsiZavody)
	if pocetZavodu < 4 {
		sql2 :=
			"SELECT z.id_zavodu,z.nazev_zavodu,z.kod_zavodu,DATE_FORMAT(z.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,z.web" +
				" FROM zavody_" + strconv.Itoa(RaceYear+1) + " z,typ_zavodu tz" +
				" WHERE" +
				" z.datum_zavodu  > CURDATE() AND" +
				" z.typ_zavodu = tz.id_typ_zavodu" +
				" ORDER BY z.datum_zavodu ASC LIMIT 0," + strconv.Itoa(4-pocetZavodu)

		results, err := db.Query(sql2)
		if err != nil {
			panic(err.Error())
		}
		//defer results.Close() přijde to tu?
		for results.Next() {
			var nejblizsiZavod NejblizsiZavod
			err = results.Scan(&nejblizsiZavod.IDZavodu, &nejblizsiZavod.NazevZavodu, &nejblizsiZavod.KodZavodu, &nejblizsiZavod.DatumZavodu, &nejblizsiZavod.Icon, &nejblizsiZavod.Web)
			nejblizsiZavody = append(nejblizsiZavody, nejblizsiZavod)
		}

	}

	json.NewEncoder(w).Encode(nejblizsiZavody)
}

func PosledniVysledky(w http.ResponseWriter, r *http.Request) {
	var posledniVysledky []NejblizsiZavod
	var posledniVysledek NejblizsiZavod

	sql1 :=
		"SELECT z.id_zavodu,z.nazev_zavodu,z.kod_zavodu,DATE_FORMAT(z.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,z.web" +
			" FROM " + SqlZavody + " z,typ_zavodu tz" +
			" WHERE" +
			" z.typ_zavodu = tz.id_typ_zavodu AND" +
			" z.zverejneni > 0 AND " +
			" z.datum_zavodu  <= CURDATE() AND" +
			" z.nove_vysledky  < 1 AND" +
			" ORDER BY z.datum_zavodu ASC LIMIT 0,4"

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close() přijde to tu?
	for results.Next() {
		err = results.Scan(&posledniVysledek.IDZavodu, &posledniVysledek.NazevZavodu, &posledniVysledek.KodZavodu, &posledniVysledek.DatumZavodu, &posledniVysledek.Icon, &posledniVysledek.Web)
		posledniVysledky = append(posledniVysledky, posledniVysledek)
	}

	pocetZavodu := len(posledniVysledky)

	if pocetZavodu < 4 {
		sql2 :=
			"SELECT z.id_zavodu,z.nazev_zavodu,z.kod_zavodu,DATE_FORMAT(z.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,z.web" +
				" FROM zavody_" + strconv.Itoa(RaceYear+1) + " z,typ_zavodu tz" +
				" WHERE" +
				" z.typ_zavodu = tz.id_typ_zavodu AND" +
				" z.zverejneni > 0 AND " +
				" z.datum_zavodu  <= CURDATE() AND" +
				" z.nove_vysledky  < 1 AND" +
				" ORDER BY z.datum_zavodu ASC LIMIT 0,4"

		results, err := db.Query(sql2)
		if err != nil {
			panic(err.Error())
		}
		//defer results.Close() přijde to tu?
		for results.Next() {
			err = results.Scan(&posledniVysledek.IDZavodu, &posledniVysledek.NazevZavodu, &posledniVysledek.KodZavodu, &posledniVysledek.DatumZavodu, &posledniVysledek.Icon, &posledniVysledek.Web)
			posledniVysledky = append(posledniVysledky, posledniVysledek)
		}

	}

	json.NewEncoder(w).Encode(posledniVysledky)
}
