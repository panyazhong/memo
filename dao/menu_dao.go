package dao

import (
	"dapan/dbx"
	"dapan/model"
)

func AddMenu(menus []model.Menu) error {
	if err := dbx.DB.Create(&menus).Error; err != nil {
		return err
	}

	return nil
}

func GetMenus(id int) ([]model.Menu, error) {
	var list []model.Menu
	if err := dbx.DB.Table("menus").Where("classify = ?", id).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
