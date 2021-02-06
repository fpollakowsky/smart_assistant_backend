package mysql

import (
	"database/sql"
	"runtime"
)

func connect() *sql.DB {
	var url string

	if runtime.GOOS == "windows" {
		url = "nethcon.com:3336"
	} else {
		url = "localhost"
	}

	db, err := sql.Open("mysql", "ecity:#BLNuo&ehd0JAPW7@tcp("+url+")/ecity")

	if err != nil {
		panic(err.Error())
	}

	return db
}
