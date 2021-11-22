package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
)

func main() {
// create client's connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error while dialling client's connection: %v", err)
		}

//	creating new client
	c := calculatorpb.NewCalculatorServiceClient(cc)

// call the functions
	//doCalculate(c)
	doPrimeNumberDecomposition(c)
}

func doCalculate(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculatorRequest{
		Number1: 3,
		Number2: 10,
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calculating the result")
	}
	log.Printf("The result of summing is: %v", res)
}

func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.Request{Result: 120}

	decomposition, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while retrieving stream of numbers: %v", err)
	}

	for {
		resp, err := decomposition.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while receiving the stream of nubmers: %v", err)
		} else {
			log.Printf("The number: %v has been decomposed into: %v", req, resp)
		}
	}
}