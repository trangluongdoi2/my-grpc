package main

import (
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg was revoked!")
	sum := float32(0)
	k := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := sum / float32(k)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float32(avg),
			})
		}

		sum += float32(req.Amount)
		k += 1
	}
}
