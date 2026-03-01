package db

import (
	"net/url"
	"strings"

	"github.com/rbrick/elections/internal/env"
	"github.com/rbrick/elections/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DefaultDbUri = "elections.db"
)

func Init(gormConfig *gorm.Config) (*gorm.DB, error) {
	rawDbUri := env.GetOrDefault("DB_URI", DefaultDbUri)
	dbMode := strings.ToLower(env.GetOrDefault("DB_MODE", "sqlite"))

	switch dbMode {
	case "sqlite":
		return initSQLite(rawDbUri, gormConfig)
	case "postgres":
		parsed, err := url.Parse(rawDbUri)

		if err != nil {
			return nil, err
		}
		return initPostgres(parsed, gormConfig)
	default:
		return nil, nil
	}
}

func initPostgres(parsed *url.URL, gormConfig *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(parsed.String()), gormConfig)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Election{}, &models.ElectionResult{}, &models.Candidate{}, &models.InternalElectionMapping{})

	return db, nil
}

func initSQLite(s string, gormConfig *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s), gormConfig)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Election{}, &models.ElectionResult{}, &models.Candidate{}, &models.InternalElectionMapping{})

	return db, nil
}
