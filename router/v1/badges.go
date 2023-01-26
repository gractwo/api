package v1

import (
	"gractwo-api/database"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserBadges(c *gin.Context) {
	type badge struct {
		Id   string
		Name string
		Desc string
		Expl string
		Img  string
		Date *time.Time
	}

	var badges []badge
	userId := c.Param("user")
	database.DB.Model(&database.User{}).Select("badges.id, badges.name, badges.desc, badges.expl, badges.img, given_badges.date").Joins("join given_badges on users.user_id = given_badges.user_id").Joins("join badges on badges.id = badges.id").Where("users.user_id = ?", userId).Scan(&badges)
	if badges == nil {
		c.JSON(404, "Not found")
	} else {
		c.JSON(200, badges)
	}
}
