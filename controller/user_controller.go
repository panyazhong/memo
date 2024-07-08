package controller

import (
	"dapan/config"
	"dapan/model"
	"dapan/service"
	"dapan/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserByOpenid(c *gin.Context) {
	username := c.Param("username")
	user, err := service.GetUserByOpenid(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func Login(c *gin.Context) {
	// var u model.UserInfo
	type LoginInfo struct {
		Code string `form:"code" binding:"required"`
	}

	var login_info LoginInfo

	err := c.ShouldBind(&login_info)

	if err != nil {
		panic(err)
	}
	fmt.Println("register login_info", login_info)

	openid, session_key, err := utils.RequestToken(config.InitWechatConfig().Appid, config.InitWechatConfig().Secret, login_info.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})

		return
	}
	_, getErr := service.GetUserByOpenid(openid)
	if getErr != nil {
		if errors.Is(getErr, gorm.ErrRecordNotFound) {
			registerUser := model.UserInfo{
				Openid:     openid,
				Sessionkey: session_key,
			}

			err := RegisterUser(registerUser)

			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
				})
				return
			} else {
				token, _ := utils.GenerToken(openid, session_key)
				c.JSON(http.StatusOK, gin.H{
					"message":  "登陆成功",
					"data":     token,
					"respCode": "000000",
				})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": getErr.Error(),
			})
			return
		}
	} else {
		token, _ := utils.GenerToken(openid, session_key)
		fmt.Println("register token 2", token)
		c.JSON(http.StatusOK, gin.H{
			"message":  "登陆成功",
			"data":     token,
			"respCode": "000000",
		})
	}

	// if u.ID == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "用户不存在",
	// 	})
	// 	return
	// }

	// if u.CheckPwd(login_info.Password) {
	// 	token, err := utils.GenerToken(u.ID)

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": token,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "密码不正确",
	// 	})
	// }

}

func RegisterUser(user model.UserInfo) error {

	if err := service.CreateUser(user); err != nil {

		return err
	}

	return nil
}
