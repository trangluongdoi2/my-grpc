package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog function was invoked %v\n", in)

	oid, err := bson.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse ID",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.Internal,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}
