package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
)

func doDeleteBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("doDeleteBlog")

	input := pb.BlogId{
		Id: id,
	}

	res, err := c.DeleteBlog(context.Background(), &input)

	if err != nil {
		log.Fatalf("Error while deleting the blog: %v", err)
	}

	log.Println("Deleting blog successfully!, %v", res)
}
