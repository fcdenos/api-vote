package sqlboiler

import (
	"database/sql"
	"log"

	"github.com/ritoon/api-vote/db"
)

type SQLboiler struct {
	db *sql.DB
}

func New(dbHost, dbName string) db.DataManager {

	var sqlB SQLboiler
	var err error

	// Open handle to database like normal
	sqlB.db, err = sql.Open("postgres", "host="+dbHost+" user=user password=pass dbname="+dbName+" sslmode=disable")
	if err != nil {
		log.Println("faile to connect", err)
		panic("failed to connect database")
	}

	// Migrate the schema
	return &sqlB
}
