package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/minecraft_bakend_web/config"
	"github.com/minecraft_bakend_web/controller"
)

func main() {
	config.LoadConfig()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://192.168.100.65:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/username/uuid/:name", controller.GetUserByUUID)
	router.POST("/login", controller.LoginUser)
	router.GET("/get_skin_data", controller.GetSkinData)
	router.GET("/get_skin_data/:uuid", controller.GetSkinByUUid)

	router.Run(":8085")

}
