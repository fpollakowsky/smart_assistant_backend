package models

type Blinder struct {
	Room    string `json:"Room"`
	Channel string `json:"Channel"`
	Value   string `json:"Value"`
}

type Light struct {
	Room    string `json:"Room"`
	Channel string `json:"Channel"`
}
