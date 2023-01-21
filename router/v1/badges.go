package v1

import (
	"gractwo-api/database"

	"github.com/gin-gonic/gin"
)

func GetUserBadges(c *gin.Context) {
	var badges []database.Badge
	userId := c.Param("user")
	database.DB.Model(&database.User{}).Select("badges.id, badges.name, badges.desc, badges.expl, badges.img").Joins("join given_badges on users.user_id = given_badges.user_id").Joins("join badges on badges.id = badges.id").Where("users.user_id = ?", userId).Scan(&badges)
	c.JSON(200, badges)
}
