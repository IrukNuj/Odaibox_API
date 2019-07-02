package main

import (
	"github.com/iruknuj/odaibox_API/db"
	"github.com/iruknuj/odaibox_API/server"
)

func main()  {
	db.Init()
	server.Init()
}