package moke

import (
	"errors"

	"github.com/ritoon/api-vote/model"
)

func (m Moke) Vote(uuid_user, uuid_proposal string) error {
	pi, ok := m.listProposal.Load(uuid_proposal)
	if !ok {
		return errors.New("proposal not found")
	}
	p := pi.(*model.Proposal)
	p.NBVote = p.NBVote + 1

	ui, ok := m.listUser.Load(uuid_user)
	if !ok {
		return errors.New("user not found")
	}
	u := ui.(*model.User)
	u.VoteDone = true

	return nil
}
