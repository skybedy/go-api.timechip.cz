package utils

import (
	"strconv"

	"go-api.timechip.cz/db"
)

func SqlShorcuts(RaceYear string) {
	db.SQLZavodySc = "zavody_" + RaceYear + " zv,"
	db.SQLZavodyBc = "zavody_" + RaceYear + " zv"
}

// Operation buď "a" pro součet, nebo "s" pro odečet
func StringIntSum(AnyString string, Addition int, Operation string) string {
	var result int
	intFromString, err := strconv.Atoi(AnyString)
	if err != nil {
	}

	switch Operation {
	case "a":
		result = intFromString + Addition
	case "s":
		result = intFromString - Addition
	}

	return strconv.Itoa(result)
}
