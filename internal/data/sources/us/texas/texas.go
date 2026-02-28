package texas

import (
	"github.com/rbrick/elections/internal/models"
)

const (
	TexasResultsURL = "https://results.texas-election.com/static/data/election"
)

/*

For Texas elections:

First, get the version from https://results.texas-election.com/static/data/Version.json

Then fetch election constants from https://results.texas-election.com/static/data/ElectionConstants_{versionNo}.json

The election constants provides the list of available elections and their IDs.

civix app: https://goelect.txelections.civixapps.com/api-ivis-system/api/v1/getFile?type=EVR_EARLYVOTING&electionId=53814&electionDate=02/27/2026

*/

type Version struct {
	VersionNumber int `json:"___versionNo"`
}

type ElectionConstants struct {
	// map of years -> elections held that year
	ElectionInfo map[string]interface{} `json:"electionInfo"`
}

type Source struct{}

func (ts *Source) GetElections() ([]models.Election, error) {
	return []models.Election{}, nil
}
