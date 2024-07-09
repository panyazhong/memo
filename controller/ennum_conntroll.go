package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FieldTypeItem struct {
	Label string `json:"label"`
	Value int16  `json:"value"`
}

type FieldTypeMapItem struct {
	Label string          `json:"label"`
	Items []FieldTypeItem `json:"items"`
}

func QueryFieldTypeMap(c *gin.Context) {
	var fieldTypeMap []FieldTypeMapItem

	fieldTypeMap = append(fieldTypeMap, FieldTypeMapItem{
		Label: "口味",
		Items: []FieldTypeItem{
			{
				Label: "不辣",
				Value: 0,
			}, {
				Label: "微辣",
				Value: 1,
			}, {
				Label: "中辣",
				Value: 2,
			},
			{
				Label: "重辣",
				Value: 3,
			},
		},
	})

	fieldTypeMap = append(fieldTypeMap, FieldTypeMapItem{
		Label: "尺寸",
		Items: []FieldTypeItem{
			{
				Label: "小份",
				Value: 4,
			},
			{
				Label: "中份",
				Value: 5,
			},
			{
				Label: "大份",
				Value: 6,
			},
		},
	})

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"respCode": "000000",
		"data":     fieldTypeMap,
	})
}
