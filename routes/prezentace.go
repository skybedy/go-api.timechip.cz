package routes

import (
	"strings"
)

//NiceString1 return...první varianta, kdy se to dělalo ze Slice a ne ze Stringu, jestli to má budoucnost, nevím
func NiceString1(str []string) string {
	return strings.ToLower(strings.ReplaceAll(strings.Join(str, " "), " ", "_"))
}

//NiceString return... prudce nedodělané, je třeba odstranit diakritiku, pořešit situaci, kdy budou dvě jméína, nebo dvě příjmení
func NiceString(str string) string {
	//x := strings.ReplaceAll(strings.Split(str, " "), " ", "_")
	//s := strings.Split(str, " ")
	//y := strings.ReplaceAll(s[0], " ", "_")
	return strings.ToLower(str)
}

/*
type Osoba struct {
	Ido       int    `json:"ido"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	BirthYear int    `json:"birthYear"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	//Team      sql.NullString `json:"team"`
	NewNumber int `json:"newNumber"`
}*/

type Osoby1 struct {
	Ido   int    `json:"ido"`
	Label string `json:"label"`
}

/*
type IdsJmenoOsoby struct {
	Ids   sql.NullInt64 `json:"ids"`
	Jmeno string        `json:"jmeno"`
}*/

/*
func PersonDetails(w http.ResponseWriter, r *http.Request) {
	var idsJmenoOsoby IdsJmenoOsoby
	ido := mux.Vars(r)["ido"]
	var count byte
	err := db.QueryRow("SELECT COUNT(*),ids FROM "+SqlZavod+" WHERE ido = ?", ido).Scan(&count, &idsJmenoOsoby.Ids)
	if err != nil {
		panic(err.Error())
	}
	if count > 0 {
		err := db.QueryRow("SELECT COUNT(*),CONCAT_WS(' ',osoby.prijmeni,osoby.jmeno) AS jmeno FROM "+SqlOsoby+" WHERE ido = ?", ido).Scan(&count, &idsJmenoOsoby.Jmeno)
		if err != nil {
			panic(err.Error())
		}
		if count > 0 {
			//Headers(w)
			json.NewEncoder(w).Encode(idsJmenoOsoby)
		}
	} else {
		var osoba Osoba
		//var posledniCislo PosledniCislo
		//	var testStruct TestStruct

		sql1 := "SELECT ido,jmeno,prijmeni,rocnik,pohlavi,psc,obec FROM osoby WHERE ido = ?"
		//fmt.Println(sql1)
		err := db.QueryRow(sql1, ido).Scan(&osoba.Ido, &osoba.Firstname, &osoba.Surname, &osoba.BirthYear, &osoba.Gender, &osoba.Country, &osoba.Team)
		if err != nil {
			panic(err.Error())
		}
		sql2 := "SELECT (posledni_cislo + 1) AS cislo_cipu FROM posledni_cislo"
		err2 := db.QueryRow(sql2).Scan(&osoba.NewNumber)
		if err2 != nil {
			panic(err.Error())
		}

		//Headers(w)
		json.NewEncoder(w).Encode(osoba)
	}

}*/

/*

func FindName(w http.ResponseWriter, r *http.Request) {
	var osoby []Osoby1
	term := strings.Split(r.URL.Query()["term"][0], " ")
	var sql1 string
	if len(term) == 1 {
		sql1 += "SELECT ido,CONCAT_WS(' ',prijmeni,jmeno,rocnik) AS label FROM " + SqlOsoby + " WHERE prijmeni_bez_diakritiky LIKE '" + NiceString(term[0]) + "%' ORDER BY prijmeni,jmeno,rocnik;"
	} else {
		sql1 += "SELECT ido,CONCAT_WS(' ',prijmeni,jmeno,rocnik) AS label FROM osoby WHERE prijmeni_bez_diakritiky LIKE '" + NiceString(term[0]) + "%' AND jmeno_bez_diakritiky LIKE '" + NiceString(term[1]) + "%' ORDER BY prijmeni,jmeno,rocnik"
	}

	results, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	//defer results.Close() přijde to tu?
	for results.Next() {
		var osoba Osoby1
		err = results.Scan(&osoba.Ido, &osoba.Label)
		osoby = append(osoby, osoba)
	}
	//Headers(w)
	json.NewEncoder(w).Encode(osoby)
}
*/
