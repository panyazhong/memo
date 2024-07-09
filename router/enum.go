package router

import (
	"dapan/controller"

	"github.com/gin-gonic/gin"
)

func SetFieldType(publicAPI *gin.RouterGroup) {
	fieldTypeMap := publicAPI.Group("/fieldType")

	fieldTypeMap.GET("/query", controller.QueryFieldTypeMap)
}
