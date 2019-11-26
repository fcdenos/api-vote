package sqlite

import "github.com/ritoon/api-vote/model"

// CRUD for Proposal
func (s Sqlite) GetProposal(uuid string) (*model.Proposal, error) { return nil, nil }
func (s Sqlite) UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {
	return nil, nil
}
func (s Sqlite) CreateProposal(u *model.Proposal) (*model.Proposal, error) { return nil, nil }
func (s Sqlite) DeleteProposal(uuid string) error                          { return nil }
