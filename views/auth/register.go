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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	u.SetPassword(u.Password)

	dbx.DB.Create(&u)

	c.JSON(http.StatusOK, gin.H{
		"message": "注册",
	})
}