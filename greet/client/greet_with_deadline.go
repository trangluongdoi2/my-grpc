package main

import (
	"context"
	"log"
	"time"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Vinh",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceed!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
