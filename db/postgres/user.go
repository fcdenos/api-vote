package postgres

import (
	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

// CRUD for User
func (s Postgres) GetUser(uuid string) (*model.User, error) {
	var u model.User
	return &u, s.db.Where("uuid = ?", uuid).Find(&u).Error
}

func (s Postgres) UpdateUser(uuid string, payload *model.User) (*model.User, error) {
	return payload, s.db.Model(&payload).Updates(model.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}).Error
}

func (s Postgres) CreateUser(u *model.User) (*model.User, error) {
	u.UUID = uuid.New().String()
	return u, s.db.Create(u).Error
}

func (s Postgres) DeleteUser(uuid string) error {
	return s.db.Where("uuid = ?", uuid).Delete(&model.User{}).Error
}

func (s Postgres) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	return &u, s.db.Where("email = ?", email).Find(&u).Error
}
