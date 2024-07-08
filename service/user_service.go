package service

import (
	"dapan/dao"
	"dapan/model"
)

func GetUserByOpenid(openid string) (*model.UserInfo, error) {
	return dao.GetUserByOpenid(openid)
}

func CreateUser(user model.UserInfo) error {
	return dao.Createuser(user)
}
