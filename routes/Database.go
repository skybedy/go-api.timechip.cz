package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)
