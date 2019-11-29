package sqlboiler

import (
	"context"

	"github.com/google/uuid"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	. "github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ritoon/api-vote/model"
	"github.com/ritoon/api-vote/models"
)

func (s SQLboiler) GetProposal(uuid string) (*model.Proposal, error) {
	var p model.Proposal
	ctx := context.Background()

	q, err := models.Proposals(Where("uuid = ?", uuid)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	p.UUID = q.UUID
	p.Title = q.Title.String
	p.Desc = q.Desc.String
	p.NBVote = q.NBVote.Int
	return &p, nil
}

func (s SQLboiler) UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {

	var p models.Proposal
	ctx := context.Background()

	p.UUID = uuid
	p.Title = null.StringFrom(payload.Title)
	p.Desc = null.StringFrom(payload.Desc)

	if _, err := p.Update(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	return payload, nil
}

func (s SQLboiler) CreateProposal(q *model.Proposal) (*model.Proposal, error) {
	var p models.Proposal
	ctx := context.Background()
	q.UUID = uuid.New().String()
	p.UUID = q.UUID
	p.Title = null.StringFrom(q.Title)
	p.Desc = null.StringFrom(q.Desc)

	if err := p.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}
	return q, nil
}

func (s SQLboiler) DeleteProposal(uuid string) error {
	var p models.Proposal
	ctx := context.Background()
	p.UUID = uuid
	_, err := p.Delete(ctx, s.db)
	return err
}
