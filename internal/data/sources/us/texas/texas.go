package texas

import (
	"github.com/rbrick/elections/internal/race"
)

const (
	BaseURL = "https://results.texas-election.com/static/data/election"
)

/*

For Texas elections:

First, get the version from https://results.texas-election.com/static/data/Version.json

Then fetch election constants from https://results.texas-election.com/static/data/ElectionConstants_{versionNo}.json

The election constants provides the list of available elections and their IDs.
*/

type Version struct {
	VersionNumber int `json:"___versionNo"`
}

type ElectionConstants struct {
	// map of years -> elections held that year
	ElectionInfo map[string]interface{} `json:"electionInfo"`
}

type Source struct{}

func (*Source) Version() {}

func (ts *Source) GetRaces() ([]race.Race, error) {
	return []race.Race{}, nil
}
