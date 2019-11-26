package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ritoon/api-vote/model"
)

type ServiceProposal struct {
	list map[string]*model.Proposal
}

func InitProposal(r *gin.Engine) {
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

	p.UUID = uuid.New().String()
	sp.list[p.UUID] = &p
	ctx.JSON(http.StatusOK, p)
}

// Get is retriving a proposal from the uuid.
func (sp *ServiceProposal) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	p, ok := sp.list[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

// Delete is deleting a proposal fron the uuid.
func (sp *ServiceProposal) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	p, ok := sp.list[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	delete(sp.list, uuid)
	ctx.JSON(http.StatusOK, p)
}

// Update is updating a proposal.
func (sp *ServiceProposal) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	p, ok := sp.list[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	var payload model.Proposal
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	p.Title = payload.Title
	p.Desc = payload.Desc

	ctx.JSON(http.StatusOK, p)
}
