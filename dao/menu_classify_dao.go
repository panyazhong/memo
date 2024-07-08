package dao

import (
	"dapan/dbx"
	"dapan/model"
)

func AddClassify(classify []model.MenuClassify) error {
	if err := dbx.DB.Create(&classify).Error; err != nil {
		return err
	}
	return nil
}

func GetMenuClassifies() ([]model.MenuClassify, error) {
	var list []model.MenuClassify
	if err := dbx.DB.Table("menu_classifies").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
