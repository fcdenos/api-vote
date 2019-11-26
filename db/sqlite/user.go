package sqlite

import "github.com/ritoon/api-vote/model"

// CRUD for User
func (s Sqlite) GetUser(uuid string) (*model.User, error)                         { return nil, nil }
func (s Sqlite) UpdateUser(uuid string, payload *model.User) (*model.User, error) { return nil, nil }
func (s Sqlite) CreateUser(u *model.User) (*model.User, error)                    { return nil, nil }
func (s Sqlite) DeleteUser(uuid string) error                                     { return nil }
