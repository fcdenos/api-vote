package service

import (
	"log"
	"net/http"

	"github.com/ritoon/api-vote/db"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/model"
)

type serviceLogin struct {
	db db.DataManager
}

func initLogin(r *gin.RouterGroup, data db.DataManager) {
	var s serviceLogin
	s.db = data
	r.POST("/login", s.login)
}

// create is creating a User.
func (sl *serviceLogin) login(ctx *gin.Context) {
	var l model.Login
	if err := ctx.BindJSON(&l); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	u, err := sl.db.GetUserByEmail(l.Email)
	log.Println("user - ", u)
	log.Println("login - ", l)
	if !u.PassIsValid(l.Pass) {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, u)
}
