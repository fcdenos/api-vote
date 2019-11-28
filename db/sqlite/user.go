package sqlite

import (
	"github.com/ritoon/api-vote/model"
)

// CRUD for User
func (s Sqlite) GetUser(uuid string) (*model.User, error) {
	var u model.User
	return &u, s.db.Where("uuid = ?", uuid).Find(&u).Error
}

func (s Sqlite) UpdateUser(uuid string, payload *model.User) (*model.User, error) {
	return payload, s.db.Model(&payload).Updates(model.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}).Error
}

func (s Sqlite) CreateUser(u *model.User) (*model.User, error) {
	return u, s.db.Create(u).Error
}

func (s Sqlite) DeleteUser(uuid string) error {
	return s.db.Where("uuid = ?", uuid).Delete(&model.User{}).Error
}

func (s Sqlite) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	return &u, s.db.Where("email = ?", email).Find(&u).Error
}
