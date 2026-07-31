package sample

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	jsonl "github.com/joaohgf/bigdatacorp/internal/inbound/jsonl"
)

// Generator creates deterministic JSONL fixtures containing valid and invalid records.
type Generator struct{}

// NewGenerator creates a sample Generator.
func NewGenerator() *Generator {
	return new(Generator)
}

// Generate writes a JSONL fixture with the requested club and player counts.
func (g *Generator) Generate(ctx context.Context, path string, clubCount, playerCount int) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create sample file: %w", err)
	}
	writer := bufio.NewWriterSize(file, 1024*1024)
	if err := g.writeInvalidSamples(writer); err != nil {
		file.Close()
		return err
	}
	encoder := json.NewEncoder(writer)
	for index := 1; index <= clubCount; index++ {
		if err := ctx.Err(); err != nil {
			file.Close()
			return err
		}
		if index%50_000 == 0 {
			if _, err := writer.WriteString("{\"club_id\":\"MALFORMED\"\n"); err != nil {
				file.Close()
				return fmt.Errorf("write malformed sample: %w", err)
			}
		}
		if err := encoder.Encode(g.newClub(index, playerCount)); err != nil {
			file.Close()
			return fmt.Errorf("encode club %d: %w", index, err)
		}
	}
	if err := writer.Flush(); err != nil {
		file.Close()
		return fmt.Errorf("flush sample file: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close sample file: %w", err)
	}
	return nil
}

// writeInvalidSamples appends malformed and incomplete records for robustness checks.
func (g *Generator) writeInvalidSamples(writer *bufio.Writer) error {
	lines := []string{
		"{\"club_id\":\"BROKEN-SYNTAX\"\n",
		"{\"club_id\":\"WRONG-TYPE\",\"championship\":\"SERIE A\",\"players\":[{\"age\":\"unknown\"}]}\n",
		"{\"club_id\":\"MISSING-FIELDS\",\"championship\":\"SERIE A\"}\n",
		"{\"club_id\":\"NULL-FIELDS\",\"name\":null,\"championship\":\"SERIE B\",\"founding_date\":null,\"colors\":null,\"players\":null}\n",
		"\n",
	}
	for _, line := range lines {
		if _, err := writer.WriteString(line); err != nil {
			return fmt.Errorf("write invalid samples: %w", err)
		}
	}
	return nil
}

// newClub creates a deterministic valid club fixture.
func (g *Generator) newClub(index, playerCount int) *jsonl.Club {
	nickname := fmt.Sprintf("Apelido %d", index)
	championship := "SERIE A"
	if index%2 == 0 {
		championship = "SERIE B"
	}
	if index%20 == 0 {
		championship = "SEM CAMPEONATO"
	}
	foundingDate := "2000-01-01"
	if index%97 == 0 {
		foundingDate = "invalid-date"
	}
	target := &jsonl.Club{
		ClubID: fmt.Sprintf("CLUB-%06d", index), Name: fmt.Sprintf("Clube %d, Futebol", index),
		Championship: championship, FoundingDate: foundingDate, City: "São Paulo", State: "SP",
		Country: "Brasil", Stadium: "Estádio \"Principal\"", President: "Presidente do Clube",
		Nickname: &nickname, Colors: []string{"azul", "branco"}, Players: make([]*jsonl.Player, 0, playerCount),
	}
	if index%101 == 0 {
		target.Nickname = nil
	}
	if index%107 == 0 {
		target.Colors = nil
	}
	for position := 1; position <= playerCount; position++ {
		target.Players = append(target.Players, g.newPlayer(index, position))
	}
	return target
}

// newPlayer creates a deterministic valid player fixture.
func (g *Generator) newPlayer(clubIndex, position int) *jsonl.Player {
	age, goals, shirt := 18+position, position-1, position
	target := &jsonl.Player{
		PlayerID: fmt.Sprintf("PLAYER-%06d-%02d", clubIndex, position),
		Name:     fmt.Sprintf("Jogador %d do Clube %d", position, clubIndex),
		Age:      &age, Goals: &goals, DebutDate: "2024-01-01", Position: "Meia", ShirtNumber: &shirt,
	}
	if clubIndex%103 == 0 && position == 1 {
		target.Age, target.Goals, target.ShirtNumber = nil, nil, nil
		target.DebutDate = "invalid-date"
	}
	return target
}
