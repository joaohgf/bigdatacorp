package jsonl

type (
	Club struct {
		Titles       *int      `json:"titles"`
		ClubID       string    `json:"club_id"`
		Name         string    `json:"name"`
		Championship string    `json:"championship"`
		City         string    `json:"city"`
		State        string    `json:"state"`
		Country      string    `json:"country"`
		Stadium      string    `json:"stadium"`
		President    string    `json:"president"`
		FoundingDate string    `json:"founding_date"`
		Nickname     *string   `json:"nickname"`
		Colors       []string  `json:"colors"`
		Players      []*Player `json:"players"`
	}
	Player struct {
		Age         *int   `json:"age"`
		Goals       *int   `json:"goals"`
		ShirtNumber *int   `json:"shirt_number"`
		MarketValue *int   `json:"market_value"`
		PlayerID    string `json:"player_id"`
		Name        string `json:"name"`
		Nationality string `json:"nationality"`
		Position    string `json:"position"`
		DebutDate   string `json:"debut_date"`
	}
)

func NewClub() *Club {
	return new(Club)
}
func NewPlayer() *Player {
	return new(Player)
}
