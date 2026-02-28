package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked!")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res := fmt.Sprintf("Hello %s", req.FirstName)
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data stream: %v", err)
		}
	}
}
