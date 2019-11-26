package moke

import (
	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/model"
)

type Moke struct {
	listUser     map[string]*model.User
	listProposal map[string]*model.Proposal
}

func New() db.DataManager {
	var m Moke
	m.listUser = make(map[string]*model.User)
	m.listProposal = make(map[string]*model.Proposal)
	return &m
}
