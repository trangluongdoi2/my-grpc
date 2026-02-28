package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Long Greet was revoked!")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
