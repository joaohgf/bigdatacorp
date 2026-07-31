package domain

import (
	"time"

	"github.com/joaohgf/bigdatacorp/internal/enum"
)

type Club struct {
	Titles       int                   `json:"title"`
	ClubID       string                `json:"club_id"`
	Name         string                `json:"name"`
	City         string                `json:"city"`
	State        string                `json:"state"`
	Country      string                `json:"country"`
	Stadium      string                `json:"stadium"`
	President    string                `json:"president"`
	Nickname     string                `json:"nickname"`
	FoundingDate time.Time             `json:"founding_date"`
	Championship enum.ChampionshipType `json:"championship"`
	Colors       []string              `json:"colors"`
	Players      []*Player             `json:"players"`
}

func NewClub() *Club {
	return new(Club)
}
