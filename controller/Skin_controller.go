package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minecraft_bakend_web/repositories"
)

func GetSkinData(c *gin.Context) {
	skins := repositories.ListSkinHandler()
	if skins == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos de la piel"})
		return
	}
	c.IndentedJSON(http.StatusOK, skins)
}

func GetSkinByUUid(c *gin.Context) {
	uuid := c.Param("uuid")
	skin := repositories.GetSkinByUuid(uuid)
	c.IndentedJSON(http.StatusOK, skin)
}
