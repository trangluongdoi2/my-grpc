package main

import (
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked!")
	var maxValue int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if req.Amount >= maxValue {
			maxValue = req.Amount
		}

		err = stream.Send(&pb.MaxResponse{
			Result: maxValue,
		})

		if err != nil {
			log.Fatalf("Error while sending data stream: %v", err)
		}
	}

}
