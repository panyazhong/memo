package main

import (
	"dapan/dbx"
	"dapan/middlewares"
	"dapan/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())

	dbx.SetMysqlDb()
	router.SetView(r)
	r.Run(":8888")
}