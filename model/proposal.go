package model

import (
	"time"

	"github.com/google/uuid"
)

type ModelDB struct {
	UUID      string     `gorm:"primary_key" json:"uuid"`
	CreatedAt time.Time  `json:"_"`
	UpdatedAt time.Time  `json:"_"`
	DeletedAt *time.Time `json:"_"`
}

func (m *ModelDB) BeforeSave() (err error) {
	m.UUID = uuid.New().String()
	return
}

// Proposal represents a proposal.
type Proposal struct {
	ModelDB
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	NBVote int    `json:"nb_vote" gorm:"column:nb_vote"`
}
