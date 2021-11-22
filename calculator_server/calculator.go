package main

// run the code from ./calculator folder !!!

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

func (s *server) PrimeNumberDecomposition(req *calculatorpb.Request, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	var k int32 = 2
	var N = req.GetResult()
	for N > 1 {
		if N % k == 0 {
			stream.Send(&calculatorpb.Response{Result: k})
			log.Printf("Number: %v has been decomposed into: %v", N, k)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
// create server
	s := grpc.NewServer()
// establish connection
	ls, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("Error while listening to the network: %v", err)
		} else {
			log.Printf("Server running...\n")
		}
// register service
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
// binding
	if err := s.Serve(ls); err != nil {
		log.Fatalf("Error while serving the connection: %v", err)
	}

}