package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked!")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}
		time.Sleep(time.Second * 1)
	}

	return &pb.GreetResponse{Result: fmt.Sprintf("Hello %s", in.FirstName)}, nil
}
