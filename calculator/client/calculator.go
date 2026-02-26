package main

import (
	"context"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{Number1: 3, Number2: 10})
	if err != nil {
		log.Fatalf("Could not calculator: %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Sum)
}
