package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model	
	Username  string `form:"username" binding:"required" gorm:"type:varchar(128)"`
	Password  string `form:"password" binding:"required" gorm:"type:varchar(128)"`
	Telephone string `form:"telephone" gorm:"type:varchar(255)"`
	Role      int    `form:"role" gorm:"type:int(64)"`
}

func (u *UserInfo) SetPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if (err != nil) {
		return
	}
	u.Password = string(hash)

	return
}

func (u *UserInfo) CheckPwd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}