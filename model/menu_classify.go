package model

import (
	"gorm.io/gorm"
)

type MenuClassify struct {
	gorm.Model
	ClassifyName string `json:"classify_name" binding:"required" gorm:"type:varchar(128);column:classify_name"`
}
