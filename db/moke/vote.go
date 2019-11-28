package moke

import (
	"errors"

	"github.com/ritoon/api-vote/model"
)

func (m Moke) Vote(uuid, _ string) error {
	pi, ok := m.listProposal.Load(uuid)
	if !ok {
		return errors.New("proposal not found")
	}
	p := pi.(*model.Proposal)
	p.NBVote = p.NBVote + 1
	return nil
}
