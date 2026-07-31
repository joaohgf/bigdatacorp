package usecase

import (
	"context"

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

func (g *Generate) Generate(ctx context.Context, sources ...*domain.Club) ([]*domain.File, error) {
	filtered := g.filter(sources...)
	generated, err := g.filer.Generate(ctx, filtered...)
	return generated, err
}

func (g *Generate) filter(sources ...*domain.Club) []*domain.Club {
	filtered := []*domain.Club{}
	for _, source := range sources {
		if source.Championship == enum.UnkownChampionshipType {
			continue
		}
		filtered = append(filtered, source)
	}
	return filtered
}
