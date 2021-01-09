package utils

import "go-api.timechip.cz/routes"

func SqlShorcuts(RaceYear string) {
	routes.SQLZavodySc = "zavody_" + RaceYear + " zv,"
	routes.SQLZavodySc = "zavody_" + RaceYear + " zv"
}
