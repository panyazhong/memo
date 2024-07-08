package router

import (
	"dapan/controller"
	"dapan/utils"

	"github.com/gin-gonic/gin"
)

func SetMenuClassify(publicAPI *gin.RouterGroup) {
	menu := publicAPI.Group("/menuClassify")
	menu.Use(utils.Auth)
	menu.POST("/addClassify", controller.AddClassify)
	menu.GET("/menuClassifies", controller.GetMenuClassifies)
}
