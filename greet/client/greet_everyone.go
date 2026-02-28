package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/trangluongdoi2/my-grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	names := []string{"Vinh", "Tri", "Tai"}

	reqs := make([]*pb.GreetRequest, 0, len(names))

	for _, name := range names {
		reqs = append(reqs, &pb.GreetRequest{
			FirstName: name,
		})
	}

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while sending data stream: %v", err)
	}

	waitc := make(chan struct{})

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

			log.Printf("GreetEveryone: %s", msg.Result)
		}
		close(waitc)
	}()

	<-waitc
}
