package model

// Proposal represents a proposal.
type Proposal struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
