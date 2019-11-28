package service

import (
	"log"
	"net/http"

	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/middlware"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/model"
)

type serviceUser struct {
	db db.DataManager
}

func initUser(r *gin.RouterGroup, data db.DataManager) {
	var s serviceUser
	s.db = data
	r.POST("/user", s.create)
	r.GET("/user/:uuid", s.get)
	jwt := middlware.NewJWT()
	r.PUT("/user/:uuid", s.update)
	r.Use(jwt.MiddlewareFunc()).DELETE("/user/:uuid", s.delete)
}

// create is creating a User.
func (sp *serviceUser) create(ctx *gin.Context) {
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

// get is retriving a User from the uuid.
func (sp *serviceUser) get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := sp.db.GetUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// delete is deleting a User fron the uuid.
func (sp *serviceUser) delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := sp.db.DeleteUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// update is updating a User.
func (sp *serviceUser) update(ctx *gin.Context) {
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

	response, _ := sp.db.GetUser(uuid)

	ctx.JSON(http.StatusOK, response)
}
