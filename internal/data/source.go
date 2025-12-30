package data

import "github.com/rbrick/elections/internal/race"

type Source interface {
	GetRaces() ([]race.Race, error)
}
