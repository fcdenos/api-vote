package model

// Proposal represents a proposal.
type Proposal struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type PayloadProposal struct {
	Proposal
}

func (pp PayloadProposal) Valid() error {
	return nil
}

type DBProposal struct {
	Proposal
}
