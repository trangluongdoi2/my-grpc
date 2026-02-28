package main

import (
	"context"
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do greet many times")

	req := &pb.GreetRequest{
		FirstName: "Vinh",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("failed to greet many times: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Greet Many Times: %s\n", msg.Result)
	}
}
