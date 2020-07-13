package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/dokshor/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculator function was invoked with %v\n", req)
	a := req.GetCalculator().GetA()
	b := req.GetCalculator().GetB()
	result := a + b
	res := &calculatorpb.CalculatorResponse{
		SumResult: strconv.FormatUint(result, 10),
	}

	return res, nil
}

func main() {
	fmt.Println("gRPC Server Version 2.0 Started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if nil != err {
		log.Fatal("Failed to listen: %w", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}
}
