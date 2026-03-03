package main

import (
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	// blogId := doCreateBlog(c)
	doReadBlog(c, "69a70036310f217b409558d5")
	// fmt.Println(blogId, "blogId...")
}
