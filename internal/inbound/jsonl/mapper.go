package jsonl

import (
	"time"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

type (
	ClubMapper struct {
		player port.ToMany[*Player, *domain.Player]
	}
	PlayerMapper struct{}
)

func NewClubMapper(
	playerMapper port.ToMany[*Player, *domain.Player],
) *ClubMapper {
	target := new(ClubMapper)
	target.player = playerMapper
	return target
}

func NewPlayerMapper() *PlayerMapper {
	return new(PlayerMapper)
}

func (m *ClubMapper) To(source *Club) *domain.Club {
	if source == nil {
		return nil
	}
	target := domain.NewClub()
	target.Titles = source.Titles
	target.ClubID = source.ClubID
	target.Name = source.Name
	target.Championship = enum.ChampionshipTypeOf(source.Championship)
	target.City = source.City
	target.State = source.State
	target.Country = source.Country
	target.Stadium = source.Stadium
	target.President = source.President
	foundingDate, err := time.Parse(time.DateOnly, source.FoundingDate)
	if err == nil {
		target.FoundingDate = &foundingDate
	}
	if source.Nickname != nil {
		target.Nickname = *source.Nickname
	}
	target.Colors = source.Colors
	target.Players = m.player.ToMany(source.Players...)
	return target
}

func (m *ClubMapper) ToMany(sources ...*Club) []*domain.Club {
	targets := []*domain.Club{}
	for _, source := range sources {
		mapped := m.To(source)
		targets = append(targets, mapped)
	}
	return targets
}

func (m *PlayerMapper) To(source *Player) *domain.Player {
	if source == nil {
		return nil
	}
	target := domain.NewPlayer()
	target.Age = source.Age
	target.Goals = source.Goals
	target.ShirtNumber = source.ShirtNumber
	target.MarketValue = source.MarketValue
	target.PlayerID = source.PlayerID
	target.Name = source.Name
	target.Nationality = source.Nationality
	target.Position = source.Position
	debutDate, err := time.Parse(time.DateOnly, source.DebutDate)
	if err == nil {
		target.DebutDate = &debutDate
	}
	return target
}

func (m *PlayerMapper) ToMany(sources ...*Player) []*domain.Player {
	targets := []*domain.Player{}
	for _, source := range sources {
		mapped := m.To(source)
		targets = append(targets, mapped)
	}
	return targets
}
