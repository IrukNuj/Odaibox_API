package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iruknuj/odaibox_API/db"
)

func main()  {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hoge")
	})

	db.Init()
	router.Run()
}