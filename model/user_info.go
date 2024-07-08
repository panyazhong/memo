package model

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Openid     string `form:"openid" binding:"required" gorm:"type:varchar(128)"`
	Sessionkey string `form:"session_key" binding:"required" gorm:"type:varchar(128)"`
	Username   string `form:"username"  gorm:"type:varchar(128)"`
	Telephone  string `form:"telephone" gorm:"type:varchar(255)"`
}

// func (u *UserInfo) SetPassword(password string) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	if err != nil {
// 		return
// 	}
// 	// u.Password = string(hash)

// 	return
// }

// func (u *UserInfo) CheckPwd(password string) bool {
// 	// err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

// 	return err == nil
// }
