package main

import (
	"context"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	req := &pb.SqrtRequest{
		Amount: n,
	}
	res, err := c.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v", err)
		}
	}

	log.Printf("doSqrt: %f", res.Result)
}
