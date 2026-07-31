package domain

import "time"

// Player represents a football player associated with a club.
type Player struct {
	Age         *int       `json:"age"`
	Goals       *int       `json:"goals"`
	ShirtNumber *int       `json:"shirt_number"`
	MarketValue *int       `json:"market_value"`
	PlayerID    string     `json:"player_id"`
	Name        string     `json:"name"`
	Nationality string     `json:"nationality"`
	Position    string     `json:"position"`
	DebutDate   *time.Time `json:"debut_date"`
}

// NewPlayer creates an empty Player.
func NewPlayer() *Player {
	return new(Player)
}
