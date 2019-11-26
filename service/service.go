package service

import "github.com/gin-gonic/gin"

// Init initialize all sub services.
func Init(r *gin.Engine) {
	initProposal(r)
	initUser(r)
}
