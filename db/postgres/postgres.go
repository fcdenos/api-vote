package postgres

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/model"
)

type Postgres struct {
	db *gorm.DB
}

func New(dbName string) db.DataManager {
	var sql Postgres
	var err error
	sql.db, err = gorm.Open("postgres", "host=db user=user password=pass dbname="+dbName+" sslmode=disable")
	if err != nil {
		log.Println("faile to connect", err)
		panic("failed to connect database")
	}

	// Migrate the schema
	sql.db.AutoMigrate(&model.User{}, &model.Proposal{})
	return &sql
}
