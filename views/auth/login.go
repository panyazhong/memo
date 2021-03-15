package auth

import (
	"dapan/dbx"
	"dapan/model"
	"dapan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserInfo
	type LoginInfo struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var login_info LoginInfo

	err := c.ShouldBind(&login_info)

	if (err != nil) {
		c.JSON(422, gin.H{
			"message": "参数错误",
		})
		return
	}

	db := dbx.DB

	db.Where("username=?", login_info.Username).First(&u)

	if (u.ID == 0) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}

	if (u.CheckPwd(login_info.Password)) {
		token, err := utils.GenerToken(u.ID)

		if (err != nil) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码不正确",
		})
	}
	
	
}
