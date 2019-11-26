package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

var list map[string]*model.User

func init() {
	list = make(map[string]*model.User)
}

func GetUser(uuid string) (*model.User, error) {
	u, ok := list[uuid]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func UpdateUser(uuid string, payload *model.User) (*model.User, error) {
	u, ok := list[uuid]
	if !ok {
		return nil, errors.New("user not found")
	}

	u.FirstName = payload.FirstName
	u.LastName = payload.LastName
	return u, nil
}

func CreateUser(u *model.User) (*model.User, error) {
	u.UUID = uuid.New().String()
	list[u.UUID] = u
	return u, nil
}

func DeleteUser(uuid string) error {
	_, ok := list[uuid]
	if !ok {
		return errors.New("user not found")
	}
	delete(list, uuid)
	return nil
}
