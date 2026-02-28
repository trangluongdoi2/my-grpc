package main

import (
	"context"
	"log"
	"time"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	nums := []int32{1, 2, 3, 4}
	reqs := make([]*pb.AvgRequest, 0, len(nums))

	for _, num := range nums {
		reqs = append(reqs, &pb.AvgRequest{
			Amount: num,
		})
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Failed to doAvg %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending the number: %d\n", req.Amount)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving the data %v", err)
	}

	log.Printf("Avg: %f", res.Result)
}
