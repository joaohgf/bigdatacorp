package csv

import (
	"fmt"
	"strings"
	"time"

	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

type (
	ClubMapper   struct{}
	PlayerMapper struct{}
)

func NewClubMapper() *ClubMapper {
	return new(ClubMapper)
}

func NewPlayerMapper() *PlayerMapper {
	return new(PlayerMapper)
}

func (m *ClubMapper) To(source *domain.Club) []string {
	target := []string{source.ClubID, source.Name, string(source.Championship), source.FoundingDate.Format(time.DateOnly),
		source.City, source.State, source.Country, source.Stadium, source.President,
		source.Nickname, strings.Join(source.Colors, "|")}
	return target
}

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

func (m *PlayerMapper) To(source *domain.Club) [][]string {
	targets := [][]string{}
	for _, player := range source.Players {
		target := []string{source.ClubID, player.PlayerID, player.Name, fmt.Sprintf("%d", player.Age),
			fmt.Sprintf("%d", player.Goals), player.DebutDate.Format(time.DateOnly), player.Position, fmt.Sprintf("%d", player.ShirtNumber)}
		targets = append(targets, target)
	}
	return targets
}

func (m *PlayerMapper) ToMany(sources ...*domain.Club) [][]string {
	targets := [][]string{
		{"Id do Clube", "Id do Jogador", "Nome", "Idade", "Gols", "Data de Estreia", "Posição", "Número da Camisa"},
	}
	for _, source := range sources {
		mapped := m.To(source)
		targets = append(targets, mapped...)
	}
	return targets
}
