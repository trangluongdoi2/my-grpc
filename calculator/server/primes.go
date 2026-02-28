package main

import (
	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	n := in.Amount
	for i := int32(2); n > 1; i++ {
		if n%i == 0 {
			if err := stream.Send(&pb.PrimesResponse{
				Result: i,
			}); err != nil {
				return err
			}
			n = n / i
		}
	}
	return nil
}
