package main

import (
	"./app/models"
	"./app/workers"
	"./config"
	"./config/admin"
	"./config/database"
	"./config/routes"
	auth "./config/siteauth"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	database.Init()
	models.Init()
	auth.Init()
	admin.Init()
	workers.Init()
	f, err := os.OpenFile("log/sql.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	models.DB.SetLogger(log.New(f, "\n", 0))
	//log.SetOutput(f)

	r := routes.Router()
	gin.SetMode(gin.DebugMode)
	println(config.Config.Port)
	log.Fatal(r.Run(fmt.Sprintf(":%v", config.Config.Port)))
}
