package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	var url string

	url = "raspberrypi"

	db, err := sql.Open("mysql", "ehome:#BLNuo&ehd0JAPW7@tcp("+url+")/ehome")

	if err != nil {
		panic(err.Error())
	}

	return db
}
