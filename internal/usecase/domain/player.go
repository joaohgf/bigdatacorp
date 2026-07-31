package domain

import "time"

type Player struct {
	Age         int       `json:"age"`
	Goals       int       `json:"goals"`
	ShirtNumber int       `json:"shirt_number"`
	MarketValue int       `json:"makert_value"`
	PlayerID    string    `json:"player_id"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
	Position    string    `json:"position"`
	DebutDate   time.Time `json:"debut_date"`
}

func NewPlayer() *Player {
	return new(Player)
}
