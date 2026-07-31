package csv

import (
	"fmt"
	"strings"
	"time"

	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

type (
	// ClubMapper converts domain clubs into CSV rows.
	ClubMapper struct{}
	// PlayerMapper converts a domain club's players into CSV rows.
	PlayerMapper struct{}
)

// NewClubMapper creates a CSV ClubMapper.
func NewClubMapper() *ClubMapper {
	return new(ClubMapper)
}

// NewPlayerMapper creates a CSV PlayerMapper.
func NewPlayerMapper() *PlayerMapper {
	return new(PlayerMapper)
}

// To maps one domain club to a CSV row.
func (m *ClubMapper) To(source *domain.Club) []string {
	target := []string{
		source.ClubID, source.Name, string(source.Championship), "",
		source.City, source.State, source.Country, source.Stadium, source.President,
		source.Nickname, strings.Join(source.Colors, "|")}
	if source.FoundingDate != nil {
		target[3] = source.FoundingDate.Format(time.DateOnly)
	}
	return target
}

// ToMany maps domain clubs to CSV rows including the header.
func (m *ClubMapper) ToMany(sources ...*domain.Club) [][]string {
	targets := [][]string{
		{"Id do Clube", "Nome", "Campeonato", "Data de Fundação", "Cidade",
			"Estado", "País", "Estádio", "Presidente", "Apelido", "Cores"},
	}
	for _, source := range sources {
		mapped := m.To(source)
		targets = append(targets, mapped)
	}
	return targets
}

// To maps a domain club's players to CSV rows.
func (m *PlayerMapper) To(source *domain.Club) [][]string {
	targets := [][]string{}
	for _, player := range source.Players {
		target := []string{source.ClubID, player.PlayerID, player.Name, "",
			"", "", player.Position, ""}
		if player.Age != nil {
			target[3] = fmt.Sprintf("%d", *player.Age)
		}
		if player.Goals != nil {
			target[4] = fmt.Sprintf("%d", *player.Goals)
		}
		if player.DebutDate != nil {
			target[5] = player.DebutDate.Format(time.DateOnly)
		}
		if player.ShirtNumber != nil {
			target[7] = fmt.Sprintf("%d", *player.ShirtNumber)
		}
		targets = append(targets, target)
	}
	return targets
}

// ToMany maps multiple clubs' players to CSV rows including the header.
func (m *PlayerMapper) ToMany(sources ...*domain.Club) [][]string {
	targets := [][]string{
		{"Id do Clube", "Id do Jogador", "Nome", "Idade",
			"Gols", "Data de Estreia", "Posição", "Número da Camisa"},
	}
	for _, source := range sources {
		mapped := m.To(source)
		targets = append(targets, mapped...)
	}
	return targets
}
