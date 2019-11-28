package sqlite

func (s Sqlite) Vote(uuid, _ string) error {
	return s.db.Exec("UPDATE proposal SET nb_vote = nb_vote + 1 WHERE proposal.uuid = " + uuid).Error
}
