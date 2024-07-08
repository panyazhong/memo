package controller

import (
	"dapan/model"
	"dapan/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func SubmitOrder(c *gin.Context) {
	type Order struct {
		Order []model.OrderDetail `json:"order"`
	}
	orderId := xid.New()

	var list Order

	err := c.ShouldBindJSON(&list)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("list is :", list)

	var sum float32

	for i := 0; i < len(list.Order); i++ {
		value := list.Order[i]
		list.Order[i].OrderId = orderId.String()
		sum += float32(value.Num) * float32(value.ActualPrice)
	}

	var order model.Order
	val, ok := c.Get("openid")
	order.OrderId = orderId.String()
	order.Amount = sum
	order.Status = 0
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "token 异常",
		})
		return
	}
	order.CustomId = val.(string)

	fmt.Println(list.Order)

	orderErr := service.AddOrder(order)
	listErr := service.AddOrderDetail(list.Order)

	if orderErr == nil && listErr == nil {
		c.JSON(http.StatusOK, gin.H{
			"message":  "下单成功",
			"respCode": "000000",
		})
		return
	}

	if orderErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": orderErr.Error(),
		})
		return
	}
	if listErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": listErr.Error(),
		})
		return
	}

}

func GetOrderList(c *gin.Context) {
	val, ok := c.Get("openid")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "token 异常",
		})
		return
	}

	order, err := service.GetOrder(val.(string))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	var orderIdList []string

	for i := 0; i < len(order); i++ {
		orderIdList = append(orderIdList, order[i].OrderId)
	}

	detail, detailErr := service.GetOrderDetail(orderIdList)
	if detailErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": detailErr.Error(),
		})
		return
	}

	type OrderResult struct {
		model.Order
		OrderDetailList []model.OrderDetail `json:"orderDetailList"`
	}
	var orderRes []OrderResult
	for i := 0; i < len(order); i++ {
		var detailList []model.OrderDetail

		orderId := order[i].OrderId
		for j := 0; j < len(detail); j++ {
			if detail[j].OrderId == orderId {
				detailList = append(detailList, detail[i])
			}
		}
		var res OrderResult
		res.OrderId = orderId
		res.CustomId = order[i].CustomId
		res.Amount = order[i].Amount
		res.Status = order[i].Status
		res.OrderDetailList = detailList

		orderRes = append(orderRes, res)

	}

	c.JSON(http.StatusOK, gin.H{
		"data":     orderRes,
		"respCode": "000000",
	})
}
