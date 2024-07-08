package service

import (
	"dapan/dao"
	"dapan/model"
)

func AddMenu(menus []model.Menu) error {
	return dao.AddMenu(menus)
}

func GetMenus(id int) ([]model.Menu, error) {
	return dao.GetMenus(id)
}
