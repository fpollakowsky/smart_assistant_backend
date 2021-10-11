package models

type Device struct {
	ID      string `json:"ID"`
	Channel string `json:"Channel"`
	Room    string `json:"Room"`
	Type    string `json:"type"`
	Value   string `json:"Value"`
	IP      string `json:"IP"`
}

type Blinder struct {
	Room    string `json:"Room"`
	Channel string `json:"Channel"`
	Value   string `json:"Value"`
}

type Light struct {
	Room    string `json:"Room"`
	Channel string `json:"Channel"`
}
