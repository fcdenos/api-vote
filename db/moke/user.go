package moke

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

func (m Moke) GetUser(uuid string) (*model.User, error) {
	u, ok := m.listUser.Load(uuid)
	if !ok {
		return nil, errors.New("user not found")
	}
	return u.(*model.User), nil
}

func (m Moke) UpdateUser(uuid string, payload *model.User) (*model.User, error) {
	ui, ok := m.listUser.Load(uuid)
	if !ok {
		return nil, errors.New("user not found")
	}
	u := ui.(*model.User)
	u.FirstName = payload.FirstName
	u.LastName = payload.LastName
	return u, nil
}

func (m Moke) CreateUser(u *model.User) (*model.User, error) {
	log.Println("CreateUser 1", u)
	u.UUID = uuid.New().String()
	u.Pass = u.HashPass(u.Pass)
	log.Println("CreateUser 2", u)
	m.listUser.Store(u.UUID, u)
	return u, nil
}

func (m Moke) DeleteUser(uuid string) error {
	_, ok := m.listUser.Load(uuid)
	if !ok {
		return errors.New("user not found")
	}
	m.listUser.Delete(uuid)
	return nil
}

func (m Moke) GetUserByEmail(email string) (*model.User, error) {
	var res *model.User
	f := func(key interface{}, value interface{}) bool {
		if value.(*model.User).Email == email {
			res = value.(*model.User)
			return false
		}
		return true
	}

	m.listUser.Range(f)

	return res, nil
}
