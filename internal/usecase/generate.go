package usecase

import (
	"context"
	"strings"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

type Generate struct {
	filer port.Filer[*domain.Club, *domain.File]
}

func NewGenerate(filer port.Filer[*domain.Club, *domain.File]) *Generate {
	target := new(Generate)
	target.filer = filer
	return target
}

func (g *Generate) Generate(ctx context.Context, sources port.Sequence[*domain.Club]) ([]*domain.File, error) {
	generated, err := g.filer.Generate(ctx, g.filter(ctx, sources))
	return generated, err
}

func (g *Generate) filter(ctx context.Context, sources port.Sequence[*domain.Club]) port.Sequence[*domain.Club] {
	return func(yield func(*domain.Club, error) bool) {
		for source, err := range sources {
			if err != nil {
				yield(nil, err)
				return
			}
			if err := ctx.Err(); err != nil {
				yield(nil, err)
				return
			}
			if !g.validClub(source) {
				continue
			}
			source.Players = g.validPlayers(source.Players)
			if !yield(source, nil) {
				return
			}
		}
	}
}

func (g *Generate) validClub(source *domain.Club) bool {
	return source != nil && strings.TrimSpace(source.ClubID) != "" &&
		source.Championship != enum.UnknownChampionshipType
}

func (g *Generate) validPlayers(sources []*domain.Player) []*domain.Player {
	targets := make([]*domain.Player, 0, len(sources))
	for _, source := range sources {
		if source == nil || strings.TrimSpace(source.PlayerID) == "" {
			continue
		}
		targets = append(targets, source)
	}
	return targets
}
