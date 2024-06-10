package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minecraft_bakend_web/models"
	"github.com/minecraft_bakend_web/repositories"
)

func LoginUser(c *gin.Context) {
	var dataLogin models.UserLoginModel
	if err := c.BindJSON(&dataLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg, uuid, err := repositories.LoginUserHandler(dataLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": msg, "user": dataLogin.User, "uuid": uuid})
}

func GetUserByUUID(c *gin.Context) {
	name := c.Param("name")
	uuid, err := repositories.FindUUIDByUsername(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"uuid":     uuid,
		"username": name,
	})
}
