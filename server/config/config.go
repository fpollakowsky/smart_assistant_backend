package config

import (
	"gorm.io/gorm"
)

var SETUP = false
var DEBUG = false

var BROKER_IP string
var BROKER_PORT string

var DSN = "ehome:#BLNuo&ehd0JAPW7@tcp(raspberrypi)/ehome?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB
