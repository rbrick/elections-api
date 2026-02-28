package data

import "github.com/rbrick/elections/internal/models"

type Source interface {
	GetElections() ([]models.Election, error)
}
