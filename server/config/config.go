package config

import (
	"gorm.io/gorm"
)

var Version = "v3"
var ENVIRONMENT *string
var IS_DEBUG *bool
var SETUP = false
var DEBUG = false
var BROKER_IP string
var BROKER_PORT string
var DSN string
var DB *gorm.DB
