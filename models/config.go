package models

import "encoding/xml"

type Config struct {
	XMLName           xml.Name    `xml:"config"`
	SQL               sql         `xml:"sql"`
	DeveloperSettings devSettings `xml:"developer_settings"`
}

type sql struct {
	Database string `xml:"database"`
	User     string `xml:"user"`
}

type mqtt struct {
	IP   string `xml:"ip"`
	Port string `xml:"port"`
}

type devSettings struct {
	Wipe bool `xml:"db_wipe"`
}
