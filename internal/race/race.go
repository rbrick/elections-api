package race

type Race struct {
	Office     string      `json:"office"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Name  string `json:"name"`
	Party string `json:"party"`
	Votes int    `json:"votes"`
}
