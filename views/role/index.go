package role

import "github.com/gin-gonic/gin"

func SetRoleRoute(publicAPI *gin.RouterGroup) {
	roleGroup := publicAPI.Group("/role")
	{
		roleGroup.GET("role_list", GetRoleList)

		roleGroup.POST("add_role", AddRole)
	}
}