package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"shome-backend/models"
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

func GetDevices() ([]models.Device, error) {
	db := connect()

	var _devices = make([]models.Device, 0)

	_rows, err := db.Query("SELECT * FROM ehome.devices")
	if err != nil {
		return nil, err
	}

	for _rows.Next() {
		var _device models.Device

		err = _rows.Scan(&_device.ID, &_device.Channel, &_device.Room, &_device.IP, &_device.Value, &_device.Type)
		if err != nil {
			return _devices, err
		}

		_devices = append(_devices, _device)
	}

	defer db.Close()
	return _devices, nil
}

func GetRooms() ([]string, error) {
	db := connect()

	var _rooms = make([]string, 0)

	_rows, err := db.Query("SELECT DISTINCT devices.room FROM ehome.devices")
	if err != nil {
		return nil, err
	}

	for _rows.Next() {
		var _room string

		err = _rows.Scan(&_room)
		if err != nil {
			return _rooms, err
		}

		_rooms = append(_rooms, _room)
	}

	defer db.Close()
	return _rooms, nil
}

func GetRoutines() ([]models.Routine, error) {
	db := connect()

	var _routines = make([]models.Routine, 0)

	_rows, err := db.Query("SELECT * FROM ehome.routines")
	if err != nil {
		return nil, err
	}

	for _rows.Next() {
		var _routine models.Routine

		err = _rows.Scan(&_routine.ID, &_routine.Device, &_routine.Payload, &_routine.Room, &_routine.Channel, &_routine.Min, &_routine.Hour, &_routine.Day, &_routine.Status)
		if err != nil {
			return _routines, err
		}

		_routines = append(_routines, _routine)
	}

	defer db.Close()
	return _routines, nil
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

func UpdateStatus(channel, room string, status float64) error {
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

func GetStatus(channel string) (int, error) {
	db := connect()
	var status int

	err := db.QueryRow("SELECT status FROM ehome.devices WHERE channel=?", channel).Scan(&status)

	if err != nil {
		return 0, err
	}

	defer db.Close()

	return status, nil
}
