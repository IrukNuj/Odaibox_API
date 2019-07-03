package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iruknuj/odaibox_API/service"
)

type Controller struct {
}

// GET /posts
func (c Controller) Index(context *gin.Context) {
	var service post.Service

	res, err := service.GetAll()
	if err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		context.JSON(200, res)
	}
}

// POST /posts
func (c Controller) Create(context *gin.Context) {
	var service post.Service
	res, err := service.CreateModel(context)

	if err != nil {
		context.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		context.JSON(201, res)
	}
}

// PUT /posts/:id
func (c Controller) Update(context *gin.Context) {
	id := context.Params.ByName("id")
	var service post.Service
	res, err := service.UpdateByID(id, context)

	if err != nil {
		context.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		context.JSON(200, res)
	}
}
