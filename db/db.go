package db

import "github.com/ritoon/api-vote/model"

type DataManager interface {
	// CRUD for Proposal
	GetProposal(uuid string) (*model.Proposal, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateProposal(uuid string, payload *model.Proposal) (*model.Proposal, error)
	CreateProposal(u *model.Proposal) (*model.Proposal, error)
	DeleteProposal(uuid string) error
	// CRUD for User
	GetUser(uuid string) (*model.User, error)
	UpdateUser(uuid string, payload *model.User) (*model.User, error)
	CreateUser(u *model.User) (*model.User, error)
	DeleteUser(uuid string) error
}
