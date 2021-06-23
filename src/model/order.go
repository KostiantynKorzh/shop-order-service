package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId     uint
	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	ID       uint
	ItemId   uint
	Quantity uint
	OrderId  uint
}
