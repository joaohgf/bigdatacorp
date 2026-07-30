package domain

type Player struct {
	Age         int
	Goals       int
	ShirtNumber int
	MarketValue int
	PlayerID    string
	Name        string
	Nationality string
	Position    string
	DebutDate   string
}

func NewPlayer() *Player {
	return new(Player)
}
