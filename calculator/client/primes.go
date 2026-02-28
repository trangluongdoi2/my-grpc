package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	fmt.Println("Primes is revoked")

	req := &pb.PrimesRequest{
		Amount: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to prime %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Result)
	}
}
