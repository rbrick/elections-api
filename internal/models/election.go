package models

type Election struct {
	// Unique ID for the election
	ID string `gorm:"column:id;primaryKey"`
	// Name of the election
	Name string `gorm:"column:name;not null"`
	// Date it was held on
	Date string `gorm:"column:date;not null"`
	// Source ID
	Source string `gorm:"column:source;not null"`

	Candidates []Candidate `gorm:"foreignKey:ElectionID;references:ID"`
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
