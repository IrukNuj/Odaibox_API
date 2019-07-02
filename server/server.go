package server

import (
	"github.com/gin-gonic/gin"
	"github.com/iruknuj/odaibox_API/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	routes := gin.Default()

	posts := routes.Group("/posts")
	{
		ctrl := post.Controller{}
		posts.GET("", ctrl.Index)
		posts.POST("", ctrl.Create)
		posts.PUT("/:id", ctrl.Update)
	}

	return routes
}