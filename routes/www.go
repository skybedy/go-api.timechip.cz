package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-api.timechip.cz/db"
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

func SqlShorcuts(RaceYear string) {
	SQLZavodySc = "zavody_" + RaceYear + " zv,"
	SQLZavodySc = "zavody_" + RaceYear + " zv"
}

func StringIntSum(AnyString string, Addition int) string {
	intFromString, err := strconv.Atoi(AnyString)
	if err != nil {
	}
	result := intFromString + Addition
	return strconv.Itoa(result)
}

//NejblizsiZavody vrací seznam 4 nejbližšéch závodů
func NejblizsiZavody(RaceYear string) []NejblizsiZavod {
	SqlShorcuts(RaceYear)

	var nejblizsiZavody []NejblizsiZavod
	//NextRaceYear, err := strconv.Atoi(RaceYear)
	//if err != nil {
	//}
	//NextRaceYear = NextRaceYear + 1

	sql1 :=
		"SELECT z.id_zavodu,z.nazev_zavodu,z.kod_zavodu,DATE_FORMAT(z.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,z.web" +
			" FROM " + SQLZavodySc + "typ_zavodu tz" +
			" WHERE" +
			" z.datum_zavodu  > CURDATE() AND" +
			" z.typ_zavodu = tz.id_typ_zavodu" +
			" ORDER BY z.datum_zavodu ASC LIMIT 0,4"

	results, err := db.Mdb.Query(sql1)
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
				" FROM zavody_" + StringIntSum(RaceYear, 1) + " z,typ_zavodu tz" +
				" WHERE" +
				" z.datum_zavodu  > CURDATE() AND" +
				" z.typ_zavodu = tz.id_typ_zavodu" +
				" ORDER BY z.datum_zavodu ASC LIMIT 0," + strconv.Itoa(4-pocetZavodu)
		fmt.Println(sql2)

		results, err := database.Db.Query(sql2)
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

	//json.NewEncoder(w).Encode(nejblizsiZavody)
	return nejblizsiZavody
}

func PosledniVysledky(RaceYear string) []NejblizsiZavod {
	var posledniVysledky []NejblizsiZavod
	var posledniVysledek NejblizsiZavod

	sql1 :=
		"SELECT z.id_zavodu,z.nazev_zavodu,z.kod_zavodu,DATE_FORMAT(z.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,z.web" +
			" FROM " + SQLZavodySc + "typ_zavodu tz" +
			" WHERE" +
			" z.typ_zavodu = tz.id_typ_zavodu AND" +
			" z.zverejneni > 0 AND " +
			" z.datum_zavodu  <= CURDATE() AND" +
			" z.nove_vysledky  < 1" +
			" ORDER BY z.datum_zavodu DESC LIMIT 0,4"

	//fmt.Println(sql1)

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
				" z.nove_vysledky  < 1" +
				" ORDER BY z.datum_zavodu DESC LIMIT 0,4"

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

	//json.NewEncoder(w).Encode(posledniVysledky)
	return posledniVysledky
}

func Neco(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	y := make(map[string][]NejblizsiZavod)
	y["nejblizsi_zavody"] = NejblizsiZavody(vars["race_year"])
	y["posledni_vysledky"] = PosledniVysledky(vars["race_year"])

	fmt.Println(y)
	json.NewEncoder(w).Encode(y)

}

type ZavodySJson struct {
	IDZavodu       string `json:"id_zavodu"`
	NazevZavodu    string `json:"nazev_zavodu"`
	KodZavodu      string `json:"kod_zavodu"`
	DatumZavodu    string `json:"datum_zavodu"`
	DenZavodu      string `json:"den_zavodu"`
	DenZavoduKonec string `json:"den_zavodu_konec"`
	MistoZavodu    string `json:"misto_zavodu"`
	Prihlasky      string `json:"prihlasky"`
	NoveVysledky   string `json:"nove_vysledky"`
	Web            string `json:"web"`
	Icon           string `json:"icon"`
}

func Zavody(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(SqlZavody)

	sql1 := "SELECT " + SqlZavody + ".id_zavodu," + SqlZavody + ".nazev_zavodu,DATE_FORMAT(" + SqlZavody + ".datum_zavodu,'%e. %c') AS datum,DATE_FORMAT(" + SqlZavody + ".datum_zavodu,'%e') AS den_zavodu,DATE_FORMAT(" + SqlZavody + ".datum_zavodu_konec,'%e. %c') AS datum_zavodu_konec," + SqlZavody + ".misto_zavodu," + SqlZavody + ".web," + SqlZavody + ".prihlasky," + SqlZavody + ".nove_vysledky,typ_zavodu.typ_zavodu," + vars["race-year"] + " AS year FROM " + SqlZavody + ",typ_zavodu WHERE " + SqlZavody + ".typ_zavodu = typ_zavodu.id_typ_zavodu AND zverejneni > 0 ORDER BY datum_zavodu,nazev_zavodu"
	fmt.Println(sql1)
	//json.NewEncoder(w).Encode(y)

}
