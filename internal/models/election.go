package models

type InternalElectionMapping struct {
	// the internal ID of the election, which is the same as the Election.ID field
	ID               string `gorm:"column:id;primaryKey;autoIncrement:true"`
	SourceElectionID string `gorm:"column:source_election_id;not null"`
	Source           string `gorm:"column:source;not null"`
	ElectionID       string `gorm:"column:election_id;not null"`
}

type Election struct {
	// Unique ID for the election
	ID string `gorm:"column:id;primaryKey"`
	// Name of the election
	Name string `gorm:"column:name;not null"`
	// Date it was held on
	Date string `gorm:"column:date;not null"`
	// Source ID
	Source string `gorm:"column:source;not null"`

	ElectionMapping InternalElectionMapping `gorm:"foreignKey:ElectionID;references:ID"`
	ElectionResult  ElectionResult          `gorm:"foreignKey:ElectionID;references:ID"`
	Candidates      []Candidate             `gorm:"foreignKey:ElectionID;references:ID"`
}

type ElectionResult struct {
	ID         string  `gorm:"column:id;primaryKey"`
	ElectionID string  `gorm:"column:election_id;not null"`
	Turnout    float64 `gorm:"column:turnout;not null"`
	TotalVotes int     `gorm:"column:total_votes;not null"`
}

type Candidate struct {
	ID         string `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name;not null"`
	Party      string `gorm:"column:party;not null"`
	Affliation string `gorm:"column:affiliation;not null"`
	Votes      int    `gorm:"column:votes;not null"`
	// if available, the number of early votes for this candidate
	EarlyVotes int    `gorm:"column:early_votes;not null"`
	ElectionID string `gorm:"column:election_id;not null"`
}
