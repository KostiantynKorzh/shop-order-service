package services

import (
	"order-service/src/db"
	"order-service/src/model"
)

func GetAllOrders() []model.Order {
	var orders []model.Order
	db.InitDb().Find(&orders)
	return orders
}

func GetOrderById(id uint) model.Order {
	var order model.Order
	db.InitDb().Find(&order, id)
	return order
}

func CreateNewOrder() model.Order {
	order := model.Order{
		UserId:     2,
		OrderItems: []model.OrderItem{{ItemId: 1, Quantity: 2}, {ItemId: 2, Quantity: 3, OrderId: 1}},
	}
	db.InitDb().Create(&order)
	return order
}
