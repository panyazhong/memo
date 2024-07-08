package controller

import (
	"dapan/model"
	"dapan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMenu(c *gin.Context) {
	type Menus struct {
		Menu []model.Menu `json:"menu"`
	}
	var menus Menus

	err := c.ShouldBindJSON(&menus)

	if err != nil {
		panic(err.Error())
	}

	addErr := service.AddMenu(menus.Menu)

	if addErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "新增成功",
		"respCode": "000000",
	})
}

func GetMenuList(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	list, resErr := service.GetMenus(id)

	if resErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "",
		"respCode": "000000",
		"data":     list,
	})

}
