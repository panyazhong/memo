package router

import (
	"dapan/views/auth"
	"dapan/views/memo"
	"dapan/views/role"

	"github.com/gin-gonic/gin"
)

func SetView(r *gin.Engine) {
	publicAPI := r.Group("/api")
	auth.SetAuthRoute(publicAPI)

	authAPI := r.Group("/api")
	// authAPI.Use(utils.Auth)
	memo.SetMemoRoute(authAPI)

	role.SetRoleRoute(authAPI)

	// publicAPI := r.Group("/api")
}