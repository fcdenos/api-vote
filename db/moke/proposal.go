package moke

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

func (m Moke) GetProposal(uuid string) (*model.Proposal, error) {
	ui, ok := m.listProposal.Load(uuid)
	log.Printf("ui is %T %v", ui, ui)

	if !ok {
		return nil, errors.New("Proposal not found")
	}
	log.Printf("%T - %v", ui, ui)
	return ui.(*model.Proposal), nil
}

func (m Moke) UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {
	ui, ok := m.listProposal.Load(uuid)
	if !ok {
		return nil, errors.New("Proposal not found")
	}
	u := ui.(*model.Proposal)
	u.Title = payload.Title
	u.Desc = payload.Desc
	return u, nil
}

func (m Moke) CreateProposal(u *model.Proposal) (*model.Proposal, error) {
	u.UUID = uuid.New().String()
	m.listProposal.Store(u.UUID, u)
	ui, _ := m.listProposal.Load(u.UUID)
	log.Printf("ui is %T %v", ui, ui)
	return u, nil
}

func (m Moke) DeleteProposal(uuid string) error {
	_, ok := m.listProposal.Load(uuid)
	if !ok {
		return errors.New("Proposal not found")
	}
	m.listProposal.Delete(uuid)
	return nil
}
