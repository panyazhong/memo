package dao

import (
	"dapan/dbx"
	"dapan/model"
)

func AddOrder(order model.Order) error {
	if err := dbx.DB.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func AddOrderDetal(detailList []model.OrderDetail) error {
	if err := dbx.DB.Create(&detailList).Error; err != nil {
		return err
	}

	return nil
}

func GetOrder(openid string) ([]model.Order, error) {
	var order []model.Order
	if err := dbx.DB.Table("orders").Where("custom_id = ?", openid).Scan(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderDetail(orderIdList []string) ([]model.OrderDetail, error) {
	var list []model.OrderDetail
	if err := dbx.DB.Table("order_details").Where("order_id IN ?", orderIdList).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
