package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/db"
)

// Init initialize all sub services.
func Init(r *gin.Engine, db db.DataManager) {
	initProposal(r, db)
	initUser(r, db)
}
