package v1

import (
	"gractwo-api/database"

	"github.com/gin-gonic/gin"
)

func GetAdminCards(c *gin.Context) {
	var AdminCardList []database.Admincard
	err := database.DB.Find(&AdminCardList).Error
	if err != nil {
		c.JSON(404, "Not found")
	}
	c.JSON(200, AdminCardList)
}
