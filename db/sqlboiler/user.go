package sqlboiler

import (
	"github.com/ritoon/api-vote/model"
)

// CRUD for User
func (s SQLboiler) GetUser(uuid string) (*model.User, error) {
	/*
		var u model.User
		return &u, s.db.Where("uuid = ?", uuid).Find(&u).Error
	*/
	return nil, nil
}

func (s SQLboiler) UpdateUser(uuid string, payload *model.User) (*model.User, error) {
	return nil, nil
}

func (s SQLboiler) CreateUser(u *model.User) (*model.User, error) {
	return nil, nil
}

func (s SQLboiler) DeleteUser(uuid string) error {
	return nil
}

func (s SQLboiler) GetUserByEmail(email string) (*model.User, error) {
	return nil, nil
}
