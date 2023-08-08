package config

import (
	"encoding/xml"
	"home-ai-backend/models"
	"home-ai-backend/server/logging"
	"io"
	"os"
)

var Config models.Config

func Importer() {
	var configName string

	switch *ENVIRONMENT {
	case "dev":
		configName = "dev-config.xml"
	case "prod":
		configName = "prod-config.xml"
	case "stage":
		configName = "stage-config.xml"
	default:
		return
	}

	xmlFile, err := os.Open(configName)
	if err != nil {
		// todo add error handling
	}

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			// todo add error handling
		}
	}(xmlFile)

	byteValue, _ := io.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &Config)
	if err != nil {
		// todo add error handling
	}

	DSN = Config.SQL.User + ":#BLNuo&ehd0JAPW7@tcp(raspberrypi)/" + Config.SQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	logging.PrintListDone("Import Config")
}
