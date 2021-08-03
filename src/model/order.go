package model

import "time"

type Order struct {
	ID         uint
	UserId     uint
	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     string
}

type OrderItem struct {
	ID       uint
	ItemId   uint
	Quantity uint
	OrderId  uint
}
