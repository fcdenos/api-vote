package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ritoon/api-vote/service"
)

func main() {
	r := gin.Default()
	service.InitProposal(r)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
