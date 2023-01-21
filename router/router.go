package router

import (
	v1 "gractwo-api/router/v1"

	"github.com/gin-gonic/gin"
)

//docs "github.com/go-project-name/docs"

func InitRouter() *gin.Engine {
	r := gin.Default()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//CORS

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/persons-of-note", v1.GetAdminCards)
		apiv1.GET("/index-images", v1.GetImages)
		apiv1.GET("/splash", v1.GetSplash)
		apiv1.GET("/badges/:user", v1.GetUserBadges)
	}
	return r
}
