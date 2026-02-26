package main

import (
	"context"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Sum: in.Number1 + in.Number2,
	}, nil
}
