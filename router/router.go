package router

import "github.com/gin-gonic/gin"

//docs "github.com/go-project-name/docs"

func InitRouter() *gin.Engine {
	r := gin.New()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		r.GET("./", )
	}

	return r
}
