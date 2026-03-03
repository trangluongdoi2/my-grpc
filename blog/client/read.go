package main

import (
	"context"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
)

func doReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("doReadBlog was invoked with %s:", id)

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happend while reading: %v", err)
	}

	log.Printf("Blog was read: %v", res)
	return res
}
