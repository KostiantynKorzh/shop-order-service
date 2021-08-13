package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "order-service/grpc/proto"
)

func main() {
	conn, err := grpc.Dial(":50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)
	in := &pb.OrderRequest{
		OrderId:     3,
		UserId:      3,
		TotalAmount: 33.33,
	}
	resp, err := client.SendOrderToPaymentService(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	log.Printf(resp.String())
}
