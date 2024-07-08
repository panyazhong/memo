package router

import (
	"dapan/controller"

	"github.com/gin-gonic/gin"
)

func setMenuRoute(groupAPI *gin.RouterGroup) {
	menu := groupAPI.Group("/menu")

	menu.POST("/add", controller.AddMenu)
	menu.GET("/getMenus", controller.GetMenuList)
}
