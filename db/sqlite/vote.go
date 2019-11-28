package sqlite

import "github.com/ritoon/api-vote/model"

func (s Sqlite) Vote(uuid_user, uuid_proposal string) error {
	user := model.User{
		ModelDB: model.ModelDB{
			UUID: uuid_user,
		},
	}
	err := s.db.Model(&user).Updates(model.User{
		VoteDone: true,
	}).Error
	if err != nil {
		return err
	}

	return s.db.Exec("UPDATE proposal SET nb_vote = nb_vote + 1 WHERE proposal.uuid = " + uuid_proposal).Error
}
