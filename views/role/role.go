package role

import (
	"dapan/dbx"
	"dapan/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetRoleList(c *gin.Context) {
	type Role struct {
		RoleName string `json:"role_name"`
		Id uint `json:"id"`
	}

	var roles []Role

	db := dbx.DB

	if err := db.Table("role_infos").Find(&roles); err.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"message": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": roles,
	})
}

func AddRole(c *gin.Context) {
	var role_info model.RoleInfo
	err := c.ShouldBind(&role_info)

	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"message": err.Error(),
		})
		return
	}

	if err := dbx.DB.Create(&role_info); err.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message": "创建成功",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"err": err.Error,
	})
	
}