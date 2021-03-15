package auth

import "github.com/gin-gonic/gin"

func SetAuthRoute(publicAPI *gin.RouterGroup) {
	authGroup := publicAPI.Group("/auth")
	{
		authGroup.POST("/register", Register)

		authGroup.POST("/login", Login)
	}
}