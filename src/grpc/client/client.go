package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "order-service/grpc/proto"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
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
	fmt.Println(in)
	resp, err := client.ProcessPayment(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	fmt.Println(resp.String())
}
