package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
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


