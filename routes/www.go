package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-api.timechip.cz/db"
	"go-api.timechip.cz/utils"
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
func NejblizsiZavody(RaceYear string) []NejblizsiZavod {
	utils.SqlShorcuts(RaceYear)
	var nejblizsiZavody []NejblizsiZavod

	sql1 :=
		"SELECT zv.id_zavodu,zv.nazev_zavodu,zv.kod_zavodu,DATE_FORMAT(zv.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,zv.web" +
			" FROM " + db.SQLZavodySc + "typ_zavodu tz" +
			" WHERE" +
			" zv.datum_zavodu  > CURDATE() AND" +
			" zv.typ_zavodu = tz.id_typ_zavodu" +
			" ORDER BY zv.datum_zavodu ASC LIMIT 0,4"

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
			"SELECT zv.id_zavodu,zv.nazev_zavodu,zv.kod_zavodu,DATE_FORMAT(zv.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,zv.web" +
				" FROM zavody_" + utils.StringIntSum(RaceYear, 1, "a") + " zv,typ_zavodu tz" +
				" WHERE" +
				" zv.datum_zavodu  > CURDATE() AND" +
				" zv.typ_zavodu = tz.id_typ_zavodu" +
				" ORDER BY zv.datum_zavodu ASC LIMIT 0," + strconv.Itoa(4-pocetZavodu)

		results, err := db.Mdb.Query(sql2)
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

	return nejblizsiZavody
}

func PosledniVysledky(RaceYear string) []NejblizsiZavod {
	var posledniVysledky []NejblizsiZavod
	var posledniVysledek NejblizsiZavod

	sql1 :=
		"SELECT zv.id_zavodu,zv.nazev_zavodu,zv.kod_zavodu,DATE_FORMAT(zv.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,zv.web" +
			" FROM " + db.SQLZavodySc + "typ_zavodu tz" +
			" WHERE" +
			" zv.typ_zavodu = tz.id_typ_zavodu AND" +
			" zv.zverejneni > 0 AND " +
			" zv.datum_zavodu  <= CURDATE() AND" +
			" zv.nove_vysledky  < 1" +
			" ORDER BY zv.datum_zavodu DESC LIMIT 0,4"
	results, err := db.Mdb.Query(sql1)
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
			"SELECT zv.id_zavodu,zv.nazev_zavodu,zv.kod_zavodu,DATE_FORMAT(zv.datum_zavodu,'%e.%c.%Y') AS datum,tz.icon,zv.web" +
				" FROM zavody_" + utils.StringIntSum(RaceYear, 1, "s") + " zv,typ_zavodu tz" +
				" WHERE" +
				" zv.typ_zavodu = tz.id_typ_zavodu AND" +
				" zv.zverejneni > 0 AND " +
				" zv.datum_zavodu  <= CURDATE() AND" +
				" zv.nove_vysledky  < 1" +
				" ORDER BY zv.datum_zavodu DESC LIMIT 0,4"

		results, err := db.Mdb.Query(sql2)
		if err != nil {
			panic(err.Error())
		}
		//defer results.Close() přijde to tu?
		for results.Next() {
			err = results.Scan(&posledniVysledek.IDZavodu, &posledniVysledek.NazevZavodu, &posledniVysledek.KodZavodu, &posledniVysledek.DatumZavodu, &posledniVysledek.Icon, &posledniVysledek.Web)
			posledniVysledky = append(posledniVysledky, posledniVysledek)
		}

	}

	return posledniVysledky
}

func Neco(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	y := make(map[string][]NejblizsiZavod)
	y["nejblizsi_zavody"] = NejblizsiZavody(vars["race-year"])
	y["posledni_vysledky"] = PosledniVysledky(vars["race-year"])
	json.NewEncoder(w).Encode(y)
}

type ZavodySJson struct {
	IDZavodu       string  `json:"id_zavodu"`
	NazevZavodu    string  `json:"nazev_zavodu"`
	KodZavodu      string  `json:"kod_zavodu"`
	DatumZavodu    string  `json:"datum_zavodu"`
	DenZavodu      string  `json:"den_zavodu"`
	DenZavoduKonec string  `json:"den_zavodu_konec,omitempty"` //jeden typ příatupu k NULL
	MistoZavodu    string  `json:"misto_zavodu"`
	Prihlasky      string  `json:"prihlasky"`
	NoveVysledky   string  `json:"nove_vysledky"`
	Web            *string `json:"web"` // druhý typ přístupu k null
	TypZavodu      string  `json:"typ_zavodu"`
	Rok            string  `json:"rok"`
}

func Zavody(w http.ResponseWriter, r *http.Request) {
	var Zavody []ZavodySJson
	var Zavod ZavodySJson

	vars := mux.Vars(r)
	utils.SqlShorcuts(vars["race-year"])

	sql1 := "SELECT " +
		"zv.id_zavodu, " +
		"zv.nazev_zavodu, " +
		"zv.kod_zavodu, " +
		"DATE_FORMAT(zv.datum_zavodu,'%e. %c') AS datum, " +
		"DATE_FORMAT(zv.datum_zavodu,'%e') AS den_zavodu, " +
		"IFNULL(DATE_FORMAT(zv.datum_zavodu_konec,'%e. %c'),'') AS den_zavodu_konec, " +
		"zv.misto_zavodu, " +
		"zv.prihlasky, " +
		"zv.nove_vysledky, " +
		"zv.web, " +
		"typ_zavodu.typ_zavodu, " +
		vars["race-year"] + " AS rok " +
		"FROM " + db.SQLZavodySc + "typ_zavodu " +
		"WHERE " +
		"zv.typ_zavodu = typ_zavodu.id_typ_zavodu AND " +
		"zverejneni > 0 " +
		"ORDER BY zv.datum_zavodu,zv.nazev_zavodu"

	results, err := db.Mdb.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close() přijde to tu?
	for results.Next() {
		err = results.Scan(&Zavod.IDZavodu, &Zavod.NazevZavodu, &Zavod.KodZavodu, &Zavod.DatumZavodu, &Zavod.DenZavodu, &Zavod.DenZavoduKonec, &Zavod.MistoZavodu, &Zavod.Prihlasky, &Zavod.NoveVysledky, &Zavod.Web, &Zavod.TypZavodu, &Zavod.Rok)
		Zavody = append(Zavody, Zavod)
	}

	json.NewEncoder(w).Encode(Zavody)

}
