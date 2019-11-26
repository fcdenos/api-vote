package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

type ServiceUser struct {
	list map[string]*model.User
}

func InitUser(r *gin.Engine) {
	var s ServiceUser
	s.list = make(map[string]*model.User)
	r.POST("/user", s.Create)
	r.GET("/user/:uuid", s.Get)
	r.DELETE("/user/:uuid", s.Delete)
	r.PUT("/user/:uuid", s.Update)
}

// Create is creating a User.
func (sp *ServiceUser) Create(ctx *gin.Context) {
	var p model.User
	if err := ctx.BindJSON(&p); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	p.UUID = uuid.New().String()
	sp.list[p.UUID] = &p
	ctx.JSON(http.StatusOK, p)
}

// Get is retriving a User from the uuid.
func (sp *ServiceUser) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	p, ok := sp.list[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

// Delete is deleting a User fron the uuid.
func (sp *ServiceUser) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, ok := sp.list[uuid]
	if !ok {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	delete(sp.list, uuid)
	ctx.JSON(http.StatusOK, nil)
}

// Update is updating a User.
func (sp *ServiceUser) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	p, ok := sp.list[uuid]
	if !ok {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	var payload model.User
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	p.FirstName = payload.FirstName
	p.LastName = payload.LastName

	ctx.JSON(http.StatusOK, p)
}
