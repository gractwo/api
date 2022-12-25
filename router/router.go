package router

import (
	v1 "gractwo-api/router/v1"

	"github.com/gin-gonic/gin"
)

//docs "github.com/go-project-name/docs"

func InitRouter() *gin.Engine {
	r := gin.New()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/admincards", v1.GetAdminCards)
		apiv1.GET("/images", v1.GetImages)
		apiv1.GET("/splash", v1.GetSplash)
	}
	return r
}
