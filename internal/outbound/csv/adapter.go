package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

type CSV struct {
	clubMapper   port.ToMany[*domain.Club, []string]
	playerMapper port.ToMany[*domain.Club, []string]
}

func NewCSV(
	clubMapper port.ToMany[*domain.Club, []string],
	playerMapper port.ToMany[*domain.Club, []string],
) *CSV {
	target := new(CSV)
	target.clubMapper = clubMapper
	target.playerMapper = playerMapper
	return target
}

func (c *CSV) Generate(ctx context.Context, sources ...*domain.Club) ([]*domain.File, error) {
	clubFile, playerFile := c.getFilesName(ctx)
	err := c.createFiles(clubFile, playerFile, sources...)
	if err != nil {
		return nil, err
	}
	files := []*domain.File{clubFile, playerFile}
	return files, nil
}

func (c *CSV) createFiles(club, player *domain.File, sources ...*domain.Club) error {
	err := c.createFile(club, c.clubMapper, sources...)
	if err != nil {
		return err
	}
	err = c.createFile(player, c.playerMapper, sources...)
	return err
}

func (c *CSV) createFile(
	file *domain.File,
	mapper port.ToMany[*domain.Club, []string],
	sources ...*domain.Club,
) error {
	filer, err := os.Create(fmt.Sprintf("%s.%s", file.Name, enum.CSVType))
	if err != nil {
		return fmt.Errorf("error genereting %s file: %w", file.Name, err)
	}
	defer filer.Close()
	writer := csv.NewWriter(filer)
	defer writer.Flush()
	mapped := mapper.ToMany(sources...)
	if err = writer.WriteAll(mapped); err != nil {
		return fmt.Errorf("error writing file %s: %w", file.Name, err)
	}
	file.Name = filer.Name()
	return nil
}

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
