package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"

	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/ritoon/api-vote/db"
	"github.com/ritoon/api-vote/db/moke"
	"github.com/ritoon/api-vote/db/postgres"
	"github.com/ritoon/api-vote/db/sqlboiler"
	"github.com/ritoon/api-vote/db/sqlite"
	"github.com/ritoon/api-vote/service"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
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
	switch viper.GetString("env") {
	case "test":
		db = moke.New()
	case "preprod":
		db = sqlite.New("test.db")
	case "prod":
		db = postgres.New("localhost", "postgres")
		//db = postgres.New(GetLocalIP(), "postgres")
	case "boiler":
		db = sqlboiler.New("localhost", "postgres")
		//db = postgres.New(GetLocalIP(), "postgres")
	}
	v1 := r.Group("/v1")
	service.Init(v1, db)
	r.Run(":" + viper.GetString("port"))
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
