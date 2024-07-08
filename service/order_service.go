package service

import (
	"dapan/dao"
	"dapan/model"
)

func AddOrder(order model.Order) error {
	return dao.AddOrder(order)
}

func AddOrderDetail(detailList []model.OrderDetail) error {
	return dao.AddOrderDetal(detailList)
}

func GetOrder(openid string) ([]model.Order, error) {
	return dao.GetOrder(openid)
}

func GetOrderDetail(orderIdList []string) ([]model.OrderDetail, error) {
	return dao.GetOrderDetail(orderIdList)
}
