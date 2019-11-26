package model

import "time"

type ModelDB struct {
	UUID      string     `gorm:"primary_key" json:"uuid"`
	CreatedAt time.Time  `json:"_"`
	UpdatedAt time.Time  `json:"_"`
	DeletedAt *time.Time `json:"_"`
}

// Proposal represents a proposal.
type Proposal struct {
	ModelDB
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
