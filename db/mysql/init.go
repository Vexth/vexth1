package mysql

import (
	"database/sql"
	"log"
)

// SQLDB
var SQLDB *sql.DB

func init() {
	var err error
	SQLDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dstest?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SQLDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
