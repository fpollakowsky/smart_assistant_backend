package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

func GetLightStatus(channel, room string) (float64, error) {
	db := connect()
	var status float64

	err := db.QueryRow("SELECT status FROM ehome.devices WHERE channel=? AND room=?", channel, room).Scan(&status)

	if err != nil {
		return 0, err
	}

	defer db.Close()
	return status, nil
}

func UpdateLightStatus(channel, room string, status float64) error {
	db := connect()

	result, err := db.Exec("UPDATE devices SET status=? WHERE channel=? AND room=?", status, channel, room)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		err := errors.New("expected to affect 1 row, affected " + strconv.FormatInt(rows, 10) + " rows instead")
		return err
	}

	defer db.Close()
	return nil
}

func GetRooms(channel string) (int, error) {
	db := connect()
	var status int

	err := db.QueryRow("SELECT status FROM ehome.devices WHERE channel=?", channel).Scan(&status)

	if err != nil {
		return 0, err
	}

	defer db.Close()

	return status, nil
}
