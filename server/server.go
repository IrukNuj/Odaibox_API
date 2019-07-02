package server

import (
	"github.com/gin-gonic/gin"
	"github.com/iruknuj/odaibox_API/controller"
)

func Init() {
	r := router()
}

func router() *gin.Engine {
	routes := gin.Default()

	posts := routes.Group("/posts")
	{
		ctrl := post.Controller{}
		posts.GET("", ctrl.Index)
		posts.POST("", ctrl.Create)
	}

	return routes
}