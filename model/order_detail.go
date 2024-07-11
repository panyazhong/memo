package model

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderId     string  `gorm:"type:varchar(64); primaryKey"`
	MenuId      int     `gorm:"type:int(20); primaryKey"`
	Price       float32 `gorm:"type:float"`
	ActualPrice float32 `gorm:"type:float"`
	Num         int     `gorm:"type:int"`
	Remark      string  `gorm:"type:varchar(64)"`
}
