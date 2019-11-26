package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/model"
)

type Sqlite struct {
	db *gorm.DB
}

func New(dbName string) db.DataManager {
	var lite Sqlite
	var err error
	lite.db, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	lite.db.AutoMigrate(&model.User{}, &model.Proposal{})
	return &lite
}
