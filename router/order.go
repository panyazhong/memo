package router

import (
	"dapan/controller"
	"dapan/utils"

	"github.com/gin-gonic/gin"
)

func setOrderRoute(publicAPI *gin.RouterGroup) {
	order := publicAPI.Group("/order")
	order.Use(utils.Auth)

	order.POST("/submit", controller.SubmitOrder)
	order.GET("/orderList", controller.GetOrderList)
}
