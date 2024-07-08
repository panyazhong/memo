package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderId  string  `gorm:"type:varchar(64); primaryKey; uniqueIndex"`
	CustomId string  `gorm:"type:varchar(64)"`
	Amount   float32 `gorm:"type:float"`
	Status   int32   `gorm:"type:int"`
}
