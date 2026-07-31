package jsonl

import (
	"os"
	"path/filepath"
	"testing"
)

type identityMapper struct{}

func (m *identityMapper) To(source *Club) *Club {
	return source
}

func TestDecodeContinuesAfterMalformedAndNullRecords(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.jsonl")
	content := "{\"club_id\":\"BEFORE\"}\n{broken\nnull\n{\"club_id\":\"AFTER\"}\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	target := NewJSONL[*Club, *Club](new(identityMapper))
	var clubs []*Club
	for club, err := range target.Decode(path) {
		if err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		clubs = append(clubs, club)
	}
	if len(clubs) != 3 || clubs[0].ClubID != "BEFORE" || clubs[1] != nil || clubs[2].ClubID != "AFTER" {
		t.Fatalf("clubs = %#v", clubs)
	}
}
