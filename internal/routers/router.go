package routers

import (
	"github.com/AxelUser/eshop/internal/middlewares"
	"github.com/AxelUser/eshop/internal/routers/handlers"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Auth
	r.POST("/auth", handlers.Auth)

	apiV1 := r.Group("/v1")
	apiV1.Use(middlewares.Jwt(make([]byte, 0)))
	{
		apiV1.GET("")
	}

	return r
}
