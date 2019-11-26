package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

var porposalList map[string]*model.Proposal

func init() {
	porposalList = make(map[string]*model.Proposal)
}

func GetProposal(uuid string) (*model.Proposal, error) {
	u, ok := porposalList[uuid]
	if !ok {
		return nil, errors.New("Proposal not found")
	}
	return u, nil
}

func UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {
	u, ok := porposalList[uuid]
	if !ok {
		return nil, errors.New("Proposal not found")
	}

	u.Title = payload.Title
	u.Desc = payload.Desc
	return u, nil
}

func CreateProposal(u *model.Proposal) (*model.Proposal, error) {
	u.UUID = uuid.New().String()
	porposalList[u.UUID] = u
	return u, nil
}

func DeleteProposal(uuid string) error {
	_, ok := porposalList[uuid]
	if !ok {
		return errors.New("Proposal not found")
	}
	delete(porposalList, uuid)
	return nil
}
