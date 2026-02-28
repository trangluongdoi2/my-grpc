package main

import (
	"context"
	"log"
	"time"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	names := []string{"Vinh", "Tri", "Tai"}

	reqs := make([]*pb.GreetRequest, 0, len(names))

	for _, name := range names {
		reqs = append(reqs, &pb.GreetRequest{
			FirstName: name,
		})
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed to doLongGreet %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending the name: %s\n", req.FirstName)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving the data %v", err)
	}

	log.Printf("LongGreet: %s", res.Result)
}
