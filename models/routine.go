package models

type Routine struct {
	ID      int    `json:"ID" gorm:"primaryKey"`
	Title   string `json:"Title"`
	Device  string `json:"Device"`
	Payload string `json:"Payload"`
	Room    string `json:"Room"`
	Channel string `json:"Channel"`
	Min     string `json:"Min"`
	Hour    string `json:"Hour"`
	Day     string `json:"Day"`
	Status  string `json:"Status"`
}
