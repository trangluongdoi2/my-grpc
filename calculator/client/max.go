package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	nums := []int32{1, 5, 3, 6, 2, 20}
	reqs := make([]*pb.MaxRequest, 0, len(nums))

	for _, num := range nums {
		reqs = append(reqs, &pb.MaxRequest{
			Amount: num,
		})
	}

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while sending data stream: %v", err)
	}

	waitc := make(chan bool)

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading data stream: %v", err)
				break
			}

			log.Printf("Max: %d", msg.Result)
		}

		waitc <- true
	}()

	<-waitc
}
