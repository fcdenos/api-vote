package service

import (
	"log"
	"net/http"

	"github.com/ritoon/api-vote/db"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/model"
)

type ServiceProposal struct {
	list map[string]*model.Proposal
}

func initProposal(r *gin.Engine) {
	var s ServiceProposal
	s.list = make(map[string]*model.Proposal)
	r.POST("/proposal", s.Create)
	r.GET("/proposal/:uuid", s.Get)
	r.DELETE("/proposal/:uuid", s.Delete)
	r.PUT("/proposal/:uuid", s.Update)
}

// Create is creating a proposal.
func (sp *ServiceProposal) Create(ctx *gin.Context) {
	var p model.Proposal
	if err := ctx.BindJSON(&p); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db.CreateProposal(&p)
	ctx.JSON(http.StatusOK, p)
}

// Get is retriving a proposal from the uuid.
func (sp *ServiceProposal) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	p, err := db.GetProposal(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

// Delete is deleting a proposal fron the uuid.
func (sp *ServiceProposal) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := db.DeleteProposal(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// Update is updating a proposal.
func (sp *ServiceProposal) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	var payload model.Proposal
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	p, _ := db.UpdateProposal(uuid, &payload)

	ctx.JSON(http.StatusOK, p)
}
