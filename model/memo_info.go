package model

import (
	"gorm.io/gorm"
)

type MemoInfo struct {
	gorm.Model
	MemoName string     `form:"memo_name" binding:"required" gorm:"type:varchar(128);column:memo_name"`
	MemoDesc string     `form:"memo_desc" binding:"required" gorm:"type:varchar(512);column:memo_desc"`
	Status   string     `form:"status" gorm:"column:status;default:'todo'"`
	Creator  uint       `gorm:"column:creator"`
	// EndTime  time.Time  `form:"end_time" binding:"required"`
}