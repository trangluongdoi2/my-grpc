package main

import (
	"context"
	"io"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func doListBlogs(c pb.BlogServiceClient) {
	log.Println("doListBlogs")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while get list blogs: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}

		log.Println("List Blog: %v", res)
	}

}
