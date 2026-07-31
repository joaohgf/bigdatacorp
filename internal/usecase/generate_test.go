package usecase

import (
	"context"
	"testing"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

// collectingFiler records clubs received from the use case.
type collectingFiler struct {
	clubs []*domain.Club
}

// Generate collects every yielded club so tests can inspect use-case output.
func (f *collectingFiler) Generate(_ context.Context, sources port.Sequence[*domain.Club]) ([]*domain.File, error) {
	for source, err := range sources {
		if err != nil {
			return nil, err
		}
		f.clubs = append(f.clubs, source)
	}
	return nil, nil
}

// TestGenerateFiltersInvalidClubsAndPlayers verifies business validation at the use-case boundary.
func TestGenerateFiltersInvalidClubsAndPlayers(t *testing.T) {
	filer := new(collectingFiler)
	target := NewGenerate(filer)
	sources := func(yield func(*domain.Club, error) bool) {
		yield(nil, nil)
		yield(&domain.Club{Championship: enum.SerieA}, nil)
		yield(&domain.Club{ClubID: "OUT", Championship: enum.UnknownChampionshipType}, nil)
		yield(&domain.Club{ClubID: "VALID", Championship: enum.SerieA, Players: []*domain.Player{
			nil, {Name: "missing id"}, {PlayerID: "P1", Name: "valid"},
		}}, nil)
	}
	if _, err := target.Generate(context.Background(), sources); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(filer.clubs) != 1 || filer.clubs[0].ClubID != "VALID" {
		t.Fatalf("clubs = %#v, want only VALID", filer.clubs)
	}
	if len(filer.clubs[0].Players) != 1 || filer.clubs[0].Players[0].PlayerID != "P1" {
		t.Fatalf("players = %#v, want only P1", filer.clubs[0].Players)
	}
}
