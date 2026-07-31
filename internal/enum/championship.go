package enum

import (
	"strings"
)

type (
	ChampionshipType string
)

const (
	SerieA                  ChampionshipType = "SERIEA"
	SerieB                  ChampionshipType = "SERIEB"
	UnknownChampionshipType                  = "UNKNOWN"
)

func ChampionshipTypeOf(source string) ChampionshipType {
	target := ChampionshipType(strings.ToUpper(strings.ReplaceAll(source, " ", "")))
	switch target {
	case SerieA, SerieB:
		return ChampionshipType(source)
	default:
		return UnknownChampionshipType
	}
}
