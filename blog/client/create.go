package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
)

func doCreateBlog(c pb.BlogServiceClient) string {
	fmt.Println("doCreateBlog()")

	blog := &pb.Blog{
		AuthorId: "Nguyen Tan Vinh",
		Title:    "Blog 1",
		Content:  "This is blog about the code...",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error while creating blog in client: %v\n", err)
	}

	log.Printf("Blog have been created!: %s\n", res.Id)

	return res.Id
}
