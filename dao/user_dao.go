package dao

import (
	"dapan/dbx"
	"dapan/model"
	"fmt"
)

func GetUserByOpenid(openid string) (*model.UserInfo, error) {
	var user model.UserInfo
	if err := dbx.DB.Where("openid = ?", openid).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func Createuser(user model.UserInfo) error {
	fmt.Println("create dao info user is:", user, &user)
	if err := dbx.DB.Debug().Create(&user).Error; err != nil {
		return err
	}
	return nil
}
