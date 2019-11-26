package service

import (
	"log"
	"net/http"

	"github.com/ritoon/api-vote/db"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/model"
)

type serviceProposal struct {
	db db.DataManager
}

func initProposal(r *gin.RouterGroup, data db.DataManager) {
	var s serviceProposal
	s.db = data
	r.POST("/proposal", s.create)
	r.GET("/proposal/:uuid", s.get)
	r.DELETE("/proposal/:uuid", s.delete)
	r.PUT("/proposal/:uuid", s.update)
}

// create is creating a proposal.
func (sp *serviceProposal) create(ctx *gin.Context) {
	var p model.Proposal
	if err := ctx.BindJSON(&p); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	sp.db.CreateProposal(&p)
	ctx.JSON(http.StatusOK, p)
}

// get is retriving a proposal from the uuid.
func (sp *serviceProposal) get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	p, err := sp.db.GetProposal(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

// delete is deleting a proposal fron the uuid.
func (sp *serviceProposal) delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := sp.db.DeleteProposal(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// update is updating a proposal.
func (sp *serviceProposal) update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	var payload model.Proposal
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	p, _ := sp.db.UpdateProposal(uuid, &payload)

	ctx.JSON(http.StatusOK, p)
}
