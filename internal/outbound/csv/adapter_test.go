package csv

import (
	"context"
	standardcsv "encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

func TestGenerateWritesRFC4180FilesWithConfiguredNames(t *testing.T) {
	t.Parallel()
	directory := t.TempDir()
	clubsPath := filepath.Join(directory, "custom-clubs.csv")
	playersPath := filepath.Join(directory, "custom-players")
	ctx := context.WithValue(context.Background(), enum.ClubFileName, clubsPath)
	ctx = context.WithValue(ctx, enum.PlayerFileName, playersPath)
	date := time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)
	age, goals, shirt := 20, 3, 10
	sources := sequence(&domain.Club{
		ClubID: "CLUB", Name: "Club, Quoted", Championship: enum.ChampionshipTypeOf("SERIE A"),
		FoundingDate: &date, President: "Person \"One\"", Colors: []string{"blue", "white"},
		Players: []*domain.Player{{PlayerID: "P1", Name: "Player, One", Age: &age, Goals: &goals,
			DebutDate: &date, Position: "Forward", ShirtNumber: &shirt}},
	})

	files, err := NewCSV(NewClubMapper(), NewPlayerMapper()).Generate(ctx, sources)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(files) != 2 || files[0].Name != clubsPath || files[1].Name != playersPath+".csv" {
		t.Fatalf("Generate() files = %#v", files)
	}
	assertCSV(t, files[0].Name, [][]string{
		{"Id do Clube", "Nome", "Campeonato", "Data de Fundação", "Cidade", "Estado", "País", "Estádio", "Presidente", "Apelido", "Cores"},
		{"CLUB", "Club, Quoted", "SERIE A", "2024-01-02", "", "", "", "", "Person \"One\"", "", "blue|white"},
	})
	assertCSV(t, files[1].Name, [][]string{
		{"Id do Clube", "Id do Jogador", "Nome", "Idade", "Gols", "Data de Estreia", "Posição", "Número da Camisa"},
		{"CLUB", "P1", "Player, One", "20", "3", "2024-01-02", "Forward", "10"},
	})
}

func TestGeneratePropagatesSequenceError(t *testing.T) {
	t.Parallel()
	directory := t.TempDir()
	ctx := context.WithValue(context.Background(), enum.ClubFileName, filepath.Join(directory, "clubs"))
	ctx = context.WithValue(ctx, enum.PlayerFileName, filepath.Join(directory, "players"))
	want := context.Canceled
	sources := func(yield func(*domain.Club, error) bool) { yield(nil, want) }

	_, err := NewCSV(NewClubMapper(), NewPlayerMapper()).Generate(ctx, sources)
	if err != want {
		t.Fatalf("Generate() error = %v, want %v", err, want)
	}
}

func sequence(clubs ...*domain.Club) port.Sequence[*domain.Club] {
	return func(yield func(*domain.Club, error) bool) {
		for _, club := range clubs {
			if !yield(club, nil) {
				return
			}
		}
	}
}

func assertCSV(t *testing.T, path string, want [][]string) {
	t.Helper()
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open CSV: %v", err)
	}
	defer file.Close()
	got, err := standardcsv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatalf("read CSV: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CSV rows = %#v, want %#v", got, want)
	}
}
