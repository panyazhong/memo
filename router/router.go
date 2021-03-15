package router

import (
	"dapan/utils"
	"dapan/views/auth"
	"dapan/views/memo"

	"github.com/gin-gonic/gin"
)

func SetView(r *gin.Engine) {
	publicAPI := r.Group("/api")
	auth.SetAuthRoute(publicAPI)

	authAPI := r.Group("/api")
	authAPI.Use(utils.Auth)
	memo.SetMemoRoute(authAPI)

	// publicAPI := r.Group("/api")
}