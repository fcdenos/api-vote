package moke

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

func (m Moke) GetUser(uuid string) (*model.User, error) {
	u, ok := m.listUser[uuid]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (m Moke) UpdateUser(uuid string, payload *model.User) (*model.User, error) {
	u, ok := m.listUser[uuid]
	if !ok {
		return nil, errors.New("user not found")
	}

	u.FirstName = payload.FirstName
	u.LastName = payload.LastName
	return u, nil
}

func (m Moke) CreateUser(u *model.User) (*model.User, error) {
	u.UUID = uuid.New().String()
	m.listUser[u.UUID] = u
	return u, nil
}

func (m Moke) DeleteUser(uuid string) error {
	_, ok := m.listUser[uuid]
	if !ok {
		return errors.New("user not found")
	}
	delete(m.listUser, uuid)
	return nil
}
