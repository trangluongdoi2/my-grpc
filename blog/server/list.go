package main

import (
	"context"
	"fmt"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(_ *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	fmt.Println("ListBlogs function wase invoked!")

	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal err: %v", err),
		)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while encoding data from MongoDB: %v", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal err: %v", err),
		)
	}

	return nil
}
