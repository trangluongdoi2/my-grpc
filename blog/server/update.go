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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Println("UpdateBlog function was invoke %v", in)

	oid, err := bson.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse ID",
		)
	}

	payload := &BlogItem{
		Title:    in.Title,
		AuthorId: in.AuthorId,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": payload},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while updating the blog: %v", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.Internal,
			"Blog not found",
		)
	}

	return &emptypb.Empty{}, nil
}
