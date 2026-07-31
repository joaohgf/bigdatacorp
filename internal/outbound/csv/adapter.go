package csv

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

// CSV streams domain clubs into separate club and player CSV files.
type CSV struct {
	clubMapper   port.To[*domain.Club, []string]
	playerMapper port.To[*domain.Club, [][]string]
}

// NewCSV creates a CSV generator with row mappers for clubs and players.
func NewCSV(
	clubMapper port.To[*domain.Club, []string],
	playerMapper port.To[*domain.Club, [][]string],
) *CSV {
	target := new(CSV)
	target.clubMapper = clubMapper
	target.playerMapper = playerMapper
	return target
}

// Generate writes the source stream to CSV files and returns their descriptors.
func (c *CSV) Generate(ctx context.Context, sources port.Sequence[*domain.Club]) ([]*domain.File, error) {
	clubFile, playerFile := c.getFilesName(ctx)
	clubFiler, err := os.Create(csvPath(clubFile.Name))
	if err != nil {
		return nil, fmt.Errorf("error generating %s file: %w", clubFile.Name, err)
	}
	defer clubFiler.Close()
	playerFiler, err := os.Create(csvPath(playerFile.Name))
	if err != nil {
		return nil, fmt.Errorf("error generating %s file: %w", playerFile.Name, err)
	}
	defer playerFiler.Close()
	clubWriter := csv.NewWriter(clubFiler)
	playerWriter := csv.NewWriter(playerFiler)
	if err := c.writeHeaders(clubWriter, playerWriter); err != nil {
		return nil, err
	}
	for source, sourceErr := range sources {
		if sourceErr != nil {
			return nil, sourceErr
		}
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if err := clubWriter.Write(c.clubMapper.To(source)); err != nil {
			return nil, fmt.Errorf("error writing club %s: %w", source.ClubID, err)
		}
		for _, player := range c.playerMapper.To(source) {
			if err := playerWriter.Write(player); err != nil {
				return nil, fmt.Errorf("error writing players from club %s: %w", source.ClubID, err)
			}
		}
	}
	clubWriter.Flush()
	playerWriter.Flush()
	if err := errors.Join(clubWriter.Error(), playerWriter.Error()); err != nil {
		return nil, fmt.Errorf("error flushing CSV files: %w", err)
	}
	clubFile.Name = clubFiler.Name()
	playerFile.Name = playerFiler.Name()
	return []*domain.File{clubFile, playerFile}, nil
}

// csvPath ensures name has a CSV extension.
func csvPath(name string) string {
	if strings.EqualFold(filepath.Ext(name), fmt.Sprintf(".%s", enum.CSVType)) {
		return name
	}
	return fmt.Sprintf("%s.%s", name, enum.CSVType)
}

// writeHeaders writes the required club and player CSV schemas.
func (c *CSV) writeHeaders(club, player *csv.Writer) error {
	clubHeader := []string{"Id do Clube", "Nome", "Campeonato", "Data de Fundação", "Cidade",
		"Estado", "País", "Estádio", "Presidente", "Apelido", "Cores"}
	playerHeader := []string{"Id do Clube", "Id do Jogador", "Nome", "Idade",
		"Gols", "Data de Estreia", "Posição", "Número da Camisa"}
	if err := club.Write(clubHeader); err != nil {
		return fmt.Errorf("error writing clubs header: %w", err)
	}
	if err := player.Write(playerHeader); err != nil {
		return fmt.Errorf("error writing players header: %w", err)
	}
	return nil
}

// getFilesName resolves default or context-provided output paths.
func (c *CSV) getFilesName(ctx context.Context) (*domain.File, *domain.File) {
	clubFile := domain.NewFile()
	clubFile.Name = string(enum.ClubFileName)
	name, ok := ctx.Value(enum.ClubFileName).(string)
	if ok {
		clubFile.Name = name
	}
	playerFile := domain.NewFile()
	playerFile.Name = string(enum.PlayerFileName)
	name, ok = ctx.Value(enum.PlayerFileName).(string)
	if ok {
		playerFile.Name = name
	}
	return clubFile, playerFile
}
