package models

type Device struct {
	ID      string `json:"ID"`
	Channel string `json:"Channel"`
	Room    string `json:"Room"`
	Type    string `json:"Type"`
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

type Routine struct {
	ID      string `json:"ID"`
	Device  string `json:"Device"`
	Payload string `json:"Payload"`
	Room    string `json:"Room"`
	Channel string `json:"Channel"`
	Min     string `json:"Min"`
	Hour    string `json:"Hour"`
	Day     string `json:"Day"`
	Status  string `json:"Status"`
}
