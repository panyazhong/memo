package controller

import (
	"dapan/model"
	"dapan/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddClassify(c *gin.Context) {
	type Classifys struct {
		Classifys []model.MenuClassify `json:"classifys"`
	}

	var classifys Classifys
	err := c.ShouldBindJSON(&classifys)
	fmt.Println("classifys", classifys)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}

	addErr := service.AddClassify(classifys.Classifys)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": addErr.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "新增成功",
		"respCode": "000000",
	})
}

func GetMenuClassifies(c *gin.Context) {
	classifies, err := service.GetMenuClassifies()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     classifies,
		"respCode": "000000",
	})
}
