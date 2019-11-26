package moke

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

func (m Moke) GetProposal(uuid string) (*model.Proposal, error) {
	u, ok := m.listProposal[uuid]
	if !ok {
		return nil, errors.New("Proposal not found")
	}
	return u, nil
}

func (m Moke) UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {
	u, ok := m.listProposal[uuid]
	if !ok {
		return nil, errors.New("Proposal not found")
	}

	u.Title = payload.Title
	u.Desc = payload.Desc
	return u, nil
}

func (m Moke) CreateProposal(u *model.Proposal) (*model.Proposal, error) {
	u.UUID = uuid.New().String()
	m.listProposal[u.UUID] = u
	return u, nil
}

func (m Moke) DeleteProposal(uuid string) error {
	_, ok := m.listProposal[uuid]
	if !ok {
		return errors.New("Proposal not found")
	}
	delete(m.listProposal, uuid)
	return nil
}
