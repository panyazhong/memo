package main

import (
	"dapan/dbx"
	"dapan/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbx.InitMysqlDb()
	router.SetView(r)
	r.Run(":8888")
}
