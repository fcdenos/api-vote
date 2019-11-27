package moke

import (
	"sync"

	"github.com/ritoon/api-vote/db"
)

type Moke struct {
	listUser     sync.Map
	listProposal sync.Map
}

func New() db.DataManager {
	var m Moke
	//m.listUser = make(sync.Map)
	//m.listProposal = make(sync.Map)
	return &m
}
