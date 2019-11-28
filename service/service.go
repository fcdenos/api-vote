package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/db"
)

// Init initialize all sub services.
func Init(r *gin.RouterGroup, db db.DataManager) {
	initLogin(r, db)
	initProposal(r, db)
	initUser(r, db)
	initVote(r, db)
}
