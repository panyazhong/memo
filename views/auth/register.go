package auth

import (
	"net/http"

	"dapan/dbx"
	"dapan/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var u model.UserInfo

	err := c.ShouldBind(&u)

	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"message": err.Error(),
		})
		return
	}

	u.SetPassword(u.Password)

	dbx.DB.Create(&u)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "注册",
	})
}