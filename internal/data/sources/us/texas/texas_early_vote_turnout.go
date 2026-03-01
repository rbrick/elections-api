package texas

import (
	"fmt"
	"io"
	"log"

	"github.com/rbrick/elections/internal/models"
)

const (
	SourceName = "texas_early_vote_turnout_source_%d"
)

type TexasTurnoutSource struct {
	civixClient CivixClient
}

func NewTexasTurnoutSource() *TexasTurnoutSource {
	return &TexasTurnoutSource{
		civixClient: NewCivixClient(),
	}
}

func (tts *TexasTurnoutSource) GetElections() ([]models.Election, error) {
	availableElections, err := tts.civixClient.GetAvailableElections()

	if err != nil {
		return nil, err
	}

	elections := []models.Election{}

	for _, election := range availableElections.Elections {
		electionData, err := tts.civixClient.GetLatestEarlyVotingData(election.ID)

		if err != nil {
			if err == io.EOF {
				log.Printf("No early voting data available for election ID %d, skipping", election.ID)
				continue
			}
			return nil, err
		}

		mappedElection := models.Election{
			ID:   fmt.Sprintf(SourceName, election.ID),
			Name: election.Name,
			Date: election.Date,
			Candidates: []models.Candidate{
				{
					Name:       election.Name,
					Votes:      electionData.TotalVotes(),
					EarlyVotes: electionData.TotalVotes(),
				},
			},
		}

		elections = append(elections, mappedElection)
	}

	return elections, nil
}
