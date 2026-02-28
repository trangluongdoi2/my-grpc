package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	amount := in.Amount

	if amount < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negavite number: %d", amount),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(amount)),
	}, nil
}
