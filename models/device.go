package models

type Device struct {
	ID      string  `json:"id" gorm:"primaryKey"`
	Channel string  `json:"channel"`
	Name    string  `json:"name"`
	Room    string  `json:"room"`
	Type    string  `json:"type"`
	Value   *string `json:"value"`
	IP      string  `json:"ip"`
}
