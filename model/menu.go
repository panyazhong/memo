package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(32);column:name;uniqueIndex"`
	Classify uint16  `gorm:"type:int(16);column:classify"`
	Price    float64 `gorm:"type:float;column:price"`
	Desc     string  `gorm:"type:varchar(128);column:desc"`
}
