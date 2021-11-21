package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"log"
	"net"
)

type server struct {}

func (s *server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculate function was invoked with %v", req)
	result := &calculatorpb.CalculatorResponse{
		Result: req.Number1 + req.Number2,
	}
	return result, nil
}

func main() {
// create server
	s := grpc.NewServer()
// establish connection
	ls, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("Error while listening to the network: %v", err)
		}
// register service
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
// binding
	if err := s.Serve(ls); err != nil {
		log.Fatalf("Error while serving the connection: %v", err)
	}
}