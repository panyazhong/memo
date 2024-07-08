package router

import (
	"dapan/controller"

	"github.com/gin-gonic/gin"
)

func SetAuthRoute(publicAPI *gin.RouterGroup) {
	authGroup := publicAPI.Group("/user")
	{
		authGroup.POST("/GetUserByOpenid", controller.GetUserByOpenid)

		authGroup.POST("/login", controller.Login)
		// authGroup.POST("/register", controller.RegisterUser)
	}
}
