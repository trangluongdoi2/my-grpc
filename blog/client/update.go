package main

import (
	"context"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
)

func doUpdateBlog(c pb.BlogServiceClient, id string) {
	log.Println("doUpdateBlog...")

	input := &pb.Blog{
		Id:       id,
		AuthorId: "Update author",
		Title:    "[UPDATE] title",
		Content:  "[UPDATE] content",
	}

	res, err := c.UpdateBlog(context.Background(), input)

	if err != nil {
		log.Fatalf("Error while updating blog: %v", err)
	}

	log.Printf("Updating blog successfully: %v", res)

}
