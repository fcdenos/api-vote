package service

import (
	"log"
	"net/http"

	"github.com/ritoon/api-vote/db"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/model"
)

type ServiceUser struct {
	db db.DataManager
}

func initUser(r *gin.Engine, data db.DataManager) {
	var s ServiceUser
	s.db = data
	r.POST("/user", s.Create)
	r.GET("/user/:uuid", s.Get)
	r.DELETE("/user/:uuid", s.Delete)
	r.PUT("/user/:uuid", s.Update)
}

// Create is creating a User.
func (sp *ServiceUser) Create(ctx *gin.Context) {
	var u model.User
	if err := ctx.BindJSON(&u); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	sp.db.CreateUser(&u)
	ctx.JSON(http.StatusOK, u)
}

// Get is retriving a User from the uuid.
func (sp *ServiceUser) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := sp.db.GetUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// Delete is deleting a User fron the uuid.
func (sp *ServiceUser) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := sp.db.DeleteUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// Update is updating a User.
func (sp *ServiceUser) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	var payload model.User
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	sp.db.UpdateUser(uuid, &payload)

	ctx.JSON(http.StatusOK, &payload)
}
