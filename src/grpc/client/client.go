package grpc

import (
	"google.golang.org/grpc"
	"log"
	pb "order-service/grpc/proto"
)

func InitGRPCClient() pb.PaymentServiceClient {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	//defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	return client
}
