package router

import (
	"github.com/gin-gonic/gin"
)

func SetView(r *gin.Engine) {
	publicAPI := r.Group("/api")
	SetAuthRoute(publicAPI)

	SetMenuClassify(publicAPI)

	setMenuRoute(publicAPI)

	setOrderRoute(publicAPI)

	// authAPI := r.Group("/api")
	// authAPI.Use(utils.Auth)
	// memo.SetMemoRoute(authAPI)

	// publicAPI := r.Group("/api")
}
