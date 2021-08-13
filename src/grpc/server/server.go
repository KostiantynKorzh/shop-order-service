package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "order-service/grpc/proto"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s server) SendOrderToPaymentService(ctx context.Context, orderRequest *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Println(orderRequest)
	return &pb.OrderResponse{Message: "Fine"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ser := &server{}
	pb.RegisterPaymentServiceServer(s, ser)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
