package postgres

import (
	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

func (s Postgres) GetProposal(uuid string) (*model.Proposal, error) {
	var p model.Proposal
	return &p, s.db.Where("uuid = ?", uuid).Find(&p).Error
}

func (s Postgres) UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {
	return payload, s.db.Model(&payload).Updates(model.Proposal{Title: payload.Title, Desc: payload.Desc}).Error
}

func (s Postgres) CreateProposal(p *model.Proposal) (*model.Proposal, error) {
	p.UUID = uuid.New().String()
	return p, s.db.Create(p).Error
}

func (s Postgres) DeleteProposal(uuid string) error {
	return s.db.Where("uuid = ?", uuid).Delete(&model.Proposal{}).Error
}
