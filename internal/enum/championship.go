package enum

import (
	"strings"
)

type (
	// ChampionshipType identifies a supported football championship.
	ChampionshipType string
)

const (
	// SerieA identifies Brazil's Série A championship.
	SerieA ChampionshipType = "SERIEA"
	// SerieB identifies Brazil's Série B championship.
	SerieB ChampionshipType = "SERIEB"
	// UnknownChampionshipType identifies unsupported championship values.
	UnknownChampionshipType = "UNKNOWN"
)

// ChampionshipTypeOf normalizes a championship name and returns its supported type.
func ChampionshipTypeOf(source string) ChampionshipType {
	target := ChampionshipType(strings.ToUpper(strings.ReplaceAll(source, " ", "")))
	switch target {
	case SerieA, SerieB:
		return ChampionshipType(source)
	default:
		return UnknownChampionshipType
	}
}
