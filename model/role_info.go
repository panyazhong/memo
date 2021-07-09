package model

import "gorm.io/gorm"

type RoleInfo struct {
	gorm.Model
	RoleName string `form:"role_name" binding:"required" gorm:"type:varchar(64)"`
}