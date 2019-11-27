package moke

import (
	"sync"

	"github.com/ritoon/api-vote/db"
)

type Moke struct {
	listUser     *sync.Map
	listProposal *sync.Map
}

func New() db.DataManager {
	var m Moke
	m.listProposal = new(sync.Map)
	m.listUser = new(sync.Map)
	return &m
}
