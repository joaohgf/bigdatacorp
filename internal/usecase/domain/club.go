package domain

type Club struct {
	Titles       int
	ClubID       string
	Name         string
	Championship string
	FoundingDate string
	City         string
	State        string
	Country      string
	Stadium      string
	President    string
	Nickname     *string
	Colors       []string
	Players      []*Player
}

func NewClub() *Club {
	return new(Club)
}
