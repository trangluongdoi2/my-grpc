package main

import (
	"fmt"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	for i := range 10 {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		if err := stream.Send(&pb.GreetResponse{
			Result: res,
		}); err != nil {
			return err
		}
	}
	return nil
}
