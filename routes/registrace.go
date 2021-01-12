package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type RaceDatazal struct {
	NazevPodzavodu string `json:"nazev_podzavodu"`
}
type EventData struct {
	NazevPodzavodu  string
	PoradiPodzavodu string
	PocetClenuTymu  string
	TypZavodnika    string
	Viditelna       string
	//Jmeno           string
	//Prijmeni        string
}

type DataCategory struct {
	IdKategorie     string
	PoradiKategorie string
	NazevKategorie  string
}

type DalsiUdaje struct {
	Tricka      string
	DalsiUdaje1 string
	//dalsi_udaje_2 string
	//dalsi_udaje_3 string
	//dalsi_udaje_4 string
	//dalsi_udaje_5 string
	//dalsi_udaje_6 string
	//dalsi_udaje_7 string
}

func dataDalsiUdaje(w http.ResponseWriter, r *http.Request) {
	var dalsiUdaje DalsiUdaje
	vars := mux.Vars(r)
	sql1 := "SELECT "
	for i := 1; i <= 1; i++ {
		sql1 += "dalsi_udaje_" + strconv.Itoa(i) + ","
	}
	sql1 += "tricka FROM prihlasky_" + vars["race_year"] + " WHERE id_zavodu = " + vars["race_id"]

	/*	err = db.QueryRow(sql1).Scan(&dalsiUdaje.DalsiUdaje1, &dalsiUdaje.Tricka)

		if err != nil {
			panic(err.Error())
		}*/
	//println(dalsiUdaje.Tricka)
	stringSlice := strings.Split(dalsiUdaje.Tricka, ",")
	//fmt.Printf("%v\n", stringSlice)
	json.NewEncoder(w).Encode(stringSlice)

}

/*
func dataKategorie(w http.ResponseWriter, r *http.Request) {
	var categories []DataCategory
	vars := mux.Vars(r)
	sql1 := "SELECT id_kategorie,poradi AS poradi_kategorie,nazev_k AS nazev_kategorie FROM kategorie_" + vars["race_year"] + " k WHERE id_zavodu = " + vars["race_id"] + " ORDER BY poradi_kategorie"
	//fmt.Fprintf(w, sql1)

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var category DataCategory
		err = results.Scan(&category.IdKategorie, &category.PoradiKategorie, &category.NazevKategorie)
		categories = append(categories, category)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
	//og.Println(categories)

}*/

/*
func dataPodzavody(w http.ResponseWriter, r *http.Request) {
	var posts []EventData

	vars := mux.Vars(r)
	var sql1 string
	sql1 = "SELECT nazev AS nazev_podzavodu,poradi_podzavodu,pocet_clenu_tymu,typ_zavodnika,viditelna FROM podzavody_" + vars["race_year"] + " WHERE id_zavodu = " + vars["race_id"]
	//fmt.Fprintf(w, sql1)

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close()
	for results.Next() {
		var post EventData
		err = results.Scan(&post.NazevPodzavodu, &post.PoradiPodzavodu, &post.PocetClenuTymu, &post.TypZavodnika, &post.Viditelna)
		posts = append(posts, post)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}
*/
