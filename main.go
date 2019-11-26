package main

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/db/moke"
	"github.com/ritoon/api-vote/service"
)

func main() {
	// config
	viper.SetConfigType("yaml")
	configFile, err := ioutil.ReadFile("./config.test.yaml")
	if err != nil {
		panic(err)
	}
	viper.ReadConfig(bytes.NewBuffer(configFile))

	// init project
	r := gin.Default()
	var db db.DataManager
	if viper.GetString("env") == "test" {
		db = moke.New()
	}
	v1 := r.Group("/v1")
	service.Init(v1, db)
	r.Run(":" + viper.GetString("port"))
}
