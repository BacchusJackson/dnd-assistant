package entities

type Note struct {
	CharacterId string `json:"character_id"`
	Body        string `json:"body"`
	Date        string `json:"date"`
}
