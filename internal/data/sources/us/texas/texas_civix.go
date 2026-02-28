package texas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CivixAPIURL = "https://goelect.txelections.civixapps.com/api-ivis-system/api/v1/getFile"
)

type CivixApiRequestType string

const (
	EVR_EARLYVOTING CivixApiRequestType = "EVR_EARLYVOTING"
	EVR_ELECTION    CivixApiRequestType = "EVR_ELECTION"
)

type CivixFileResponse struct {
	Upload string `json:"upload"`
}

func (cfr *CivixFileResponse) Decode(to interface{}) error {
	rawJson, err := base64.StdEncoding.DecodeString(cfr.Upload)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawJson, to)
}

type CivixEarlyVotingDate struct {
	Date string `json:"date"`
	ID   int    `json:"date_turnout_id"`
}

type CivixAvailableElection struct {
	ID               int                    `json:"id"`
	Type             string                 `json:"type"`
	Date             string                 `json:"election_date"`
	Name             string                 `json:"election_name"`
	Certified        bool                   `json:"certified"`
	EarlyVotingDates []CivixEarlyVotingDate `json:"early_voting_dates"`
}

func (e CivixAvailableElection) LatestDate() string {
	if len(e.EarlyVotingDates) == 0 {
		return ""
	}
	return e.EarlyVotingDates[len(e.EarlyVotingDates)-1].Date
}

type CivixEarlyVotingCountyData struct {
	Name                string `json:"name"`
	RegisteredVoters    int    `json:"registered_voters"`
	InPersonVotesOnDate int    `json:"in_person_votes_on_date"`
	TotalInPersonVotes  int    `json:"total_in_person_votes_for_election"`
	MailInVotes         int    `json:"total_mail_votes_for_election"`
}

type CivixEarlyVotingElectionData struct {
	ID              int                          `json:"election_id"`
	Type            string                       `json:"election_type"`
	EarlyVotingDate string                       `json:"early_voting_date"`
	UpdatedAt       string                       `json:"date_updated"`
	Counties        []CivixEarlyVotingCountyData `json:"turnout_by_county"`
}

type CivixEarlyVotingResponse struct {
	Elections []CivixAvailableElection `json:"elections"`
}

type CivixClient interface {
	GetAvailableElections() (*CivixEarlyVotingResponse, error)
	GetEarlyVotingData(electionID int, electionDate string) (*CivixEarlyVotingElectionData, error)
	GetLatestEarlyVotingData(electionID int) (*CivixEarlyVotingElectionData, error)
}

type client struct {
	httpClient *http.Client
}

func (c *client) makeCivixRequest(url string) (*CivixFileResponse, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var civixResp CivixFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&civixResp); err != nil {
		return nil, err
	}
	return &civixResp, nil
}

func (c *client) GetAvailableElections() (*CivixEarlyVotingResponse, error) {
	apiUrl := fmt.Sprintf("%s?type=%s", CivixAPIURL, EVR_ELECTION)

	civixResp, err := c.makeCivixRequest(apiUrl)
	if err != nil {
		return nil, err
	}

	var electionsResp CivixEarlyVotingResponse
	if err := civixResp.Decode(&electionsResp); err != nil {
		return nil, err
	}

	return &electionsResp, nil
}

func (c *client) GetEarlyVotingData(electionID int, electionDate string) (*CivixEarlyVotingElectionData, error) {
	apiUrl := fmt.Sprintf("%s?type=%s&electionId=%d&electionDate=%s", CivixAPIURL, EVR_EARLYVOTING, electionID, electionDate)

	civixResp, err := c.makeCivixRequest(apiUrl)
	if err != nil {
		return nil, err
	}

	var electionData CivixEarlyVotingElectionData
	if err := civixResp.Decode(&electionData); err != nil {
		return nil, err
	}

	return &electionData, nil
}

func (c *client) GetLatestEarlyVotingData(electionID int) (*CivixEarlyVotingElectionData, error) {
	availableElections, err := c.GetAvailableElections()

	if err != nil {
		return nil, err
	}

	for _, election := range availableElections.Elections {
		if election.ID == electionID {
			return c.GetEarlyVotingData(electionID, election.LatestDate())
		}
	}

	return nil, fmt.Errorf("election with ID %d not found", electionID)
}

func NewCivixClient() CivixClient {
	return &client{
		httpClient: &http.Client{},
	}
}
