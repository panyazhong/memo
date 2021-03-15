package memo

import "github.com/gin-gonic/gin"

func SetMemoRoute(publicAPI *gin.RouterGroup) {
	memoGroup := publicAPI.Group("/memo")

	{
		memoGroup.POST("/add_memo", AddMemo)

		memoGroup.PUT("/update_memo/:id", UpdateMemo)

		memoGroup.DELETE("/delete_memo/:id", DeleteMemo)

		memoGroup.GET("/get_memo", GetMemo)
	}
}