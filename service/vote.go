package service

import (
	"log"
	"net/http"

	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/middlware"

	"github.com/gin-gonic/gin"
)

type serviceVote struct {
	db db.DataManager
}

func initVote(r *gin.RouterGroup, data db.DataManager) {
	var s serviceVote
	s.db = data
	jwt := middlware.NewJWT()
	r.Use(jwt.MiddlewareFunc()).POST("/vote/:uuid_proposal", s.vote)
}

func (s serviceVote) vote(ctx *gin.Context) {
	proposalUUID := ctx.Param("uuid_proposal")
	userUUID := ctx.GetString("jwt_user_uuid")
	u, err := s.db.GetUser(userUUID)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if u.VoteDone {
		log.Println(err)
		ctx.JSON(http.StatusNotAcceptable, nil)
		return
	}

	if err := s.db.Vote(userUUID, proposalUUID); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}
