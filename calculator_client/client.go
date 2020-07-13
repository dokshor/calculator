package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dokshor/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if nil != err {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			A: 3,
			B: 10,
		},
	}

	res, err := client.Calculator(context.Background(), req)

	if nil != err {
		log.Fatalf("There is an error invoking the remote call %v", err)
	}

	log.Printf("Response from Calculator: %v", res.SumResult)

}
