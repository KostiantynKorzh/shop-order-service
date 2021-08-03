package model

type OrderPayment struct {
	UserId      uint    `json:"userId"`
	OrderId     uint    `json:"orderId"`
	TotalAmount float64 `json:"totalAmount"`
}
