package service

import (
	"dapan/dao"
	"dapan/model"
)

func AddClassify(classifys []model.MenuClassify) error {
	return dao.AddClassify(classifys)
}

func GetMenuClassifies() ([]model.MenuClassify, error) {
	return dao.GetMenuClassifies()
}
