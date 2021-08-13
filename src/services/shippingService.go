package services

import (
	"context"
	"fmt"
	"log"
	"order-service/db"
	"order-service/grpc/client"
	pb "order-service/grpc/proto"
	"order-service/model"
)

func SendOrderToPaymentService(order model.Order) {
	client := grpc.InitGRPCClient()

	var orderItems []model.OrderItem
	db.Db.Where("order_id = ?", order.ID).Find(&orderItems)
	var totalAmount float64
	for i, order := range FullOrder {
		price := FullOrder[i].Price
		temp := float64(order.Quantity) * price
		totalAmount += temp // TODO Add price to order object
	}

	payment := &pb.OrderRequest{
		UserId:      uint64(order.UserId),
		OrderId:     uint64(order.ID),
		TotalAmount: 12.99,
	}

	fmt.Println(payment)
	resp, err := client.ProcessPayment(context.Background(), payment)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	db.Db.Model(&order).Updates(model.Order{Status: resp.Message})
}
