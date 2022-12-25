package v1

import (
	"gractwo-api/database"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func GetSplash(c *gin.Context) {
	var SplashList []database.Splash
	err := database.DB.Find(&SplashList).Error
	if err != nil {
		c.JSON(404, "Not found")
	}
	c.JSON(200, SplashList[rand.Intn(len(SplashList)-1-0+1)+0])
}
