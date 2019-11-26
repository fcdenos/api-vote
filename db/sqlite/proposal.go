package sqlite

import "github.com/ritoon/api-vote/model"

func (s Sqlite) GetProposal(uuid string) (*model.Proposal, error) {
	var p model.Proposal
	return &p, s.db.Where("uuid = ?", uuid).Find(&p).Error
}

func (s Sqlite) UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error) {
	return payload, s.db.Model(&payload).Updates(model.Proposal{Title: payload.Title, Desc: payload.Desc}).Error
}

func (s Sqlite) CreateProposal(p *model.Proposal) (*model.Proposal, error) {
	return p, s.db.Create(p).Error
}

func (s Sqlite) DeleteProposal(uuid string) error {
	return s.db.Where("uuid = ?", uuid).Delete(&model.Proposal{}).Error
}
