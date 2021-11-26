package main

import (
	"context"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
// create client's connection
	certFile := "ssl/ca.crt"
	creds, errCreds := credentials.NewClientTLSFromFile(certFile, "")
		if errCreds != nil {
			status.FromError(errCreds)
		}

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("Error while dialling client's connection: %v", err)
		}

//	creating new client
	c := calculatorpb.NewCalculatorServiceClient(cc)

// call the functions
	 doCalculate(c)
	// doPrimeNumberDecomposition(c)
	// doAverageNumber(c)
	// doFindMaximum(c)
	doSquareRoot(c)

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

func doAverageNumber(c calculatorpb.CalculatorServiceClient) {
	request := []*calculatorpb.AverageNumberRequest{
		&calculatorpb.AverageNumberRequest{
			Number: 1,
		},
		&calculatorpb.AverageNumberRequest{
			Number: 2,
		},
		&calculatorpb.AverageNumberRequest{
			Number: 3,
		},
		&calculatorpb.AverageNumberRequest{
			Number: 4,
		},
	}

	stream, err := c.AverageNumber(context.Background())
		if err != nil {
			log.Fatalf("Error while receiving the stream: %v", err)
		}
	for _, k := range request {
		log.Printf("Sending the number: %v", k.GetNumber())
		time.Sleep(1000 * time.Millisecond)
		err := stream.Send(k)
			if err != nil {
				log.Fatalf("Error while sending AverageNumber client stream: %v", err)
			}
	}
	response, err := stream.CloseAndRecv()
	log.Printf("The average number is %v", response.GetNumber())
}

func doFindMaximum(c calculatorpb.CalculatorServiceClient) {
	numbers := []*calculatorpb.FindMaximumRequest{
		&calculatorpb.FindMaximumRequest{
			Number: 1,
		}, &calculatorpb.FindMaximumRequest{
			Number: 6,
		},&calculatorpb.FindMaximumRequest{
			Number: 2,
		},&calculatorpb.FindMaximumRequest{
			Number: 3,
		},&calculatorpb.FindMaximumRequest{
			Number: 17,
		},
	}
	stream, err := c.FindMaximum(context.Background())
		if err != nil {
			log.Fatalf("Error while setting client stream: %v", err)
		}
	waitc := make(chan struct{})
	go func() {
		for _, j := range numbers {
			err := stream.Send(j)
			log.Printf("Sending the number: %v", j.Number)
			time.Sleep(2000 * time.Millisecond)
				if err != nil {
					log.Fatalf("Error while sending the client stream to the server: %v", err)
				}
		}
		stream.CloseSend()
			if err != nil {
			log.Fatalf("Error while closing the client stream: %v", err)
			}
	}()

	go func() {
		for {
			response, err := stream.Recv()
				if err == io.EOF {
					close(waitc)
					break
				} else if err != nil {
					log.Fatalf("Error while receiving server stream: %v", err)
				} else {
					log.Printf("The biggest number so far received is: %v", response.GetNumber())
				}
		}
	}()
<-waitc
}

func doSquareRoot(c calculatorpb.CalculatorServiceClient) {
	var number int32 = 9

	response, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: number})
		if err != nil {
			status, ok := status.FromError(err)
				if ok == true {
					log.Printf("Error message from server: %v\n", status.Message())
					log.Println(status.Code())
					return
				} else {
					log.Fatalf("Some big trouble: %v", err)
					return
				}
		}
	log.Printf("Square root of number %v is equal to %v", number, response.GetNumber())
}






