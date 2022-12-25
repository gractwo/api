package v1

import (
	"gractwo-api/database"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context){
	var ImageList []database.Image
	err :=database.DB.Find(&ImageList).Error
	if err != nil {
		c.JSON(404, "Not found")
	}
	c.JSON(200, ImageList)
}