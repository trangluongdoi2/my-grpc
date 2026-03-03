package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with: %v\n", in)

	oid, err := bson.ObjectIDFromHex(in.Id)
	// oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		log.Printf("Failed to parse ObjectID from hex: %v\n", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	log.Printf("Looking for blog with _id: %v\n", oid)

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		fmt.Println(err.Error(), "errrrr")
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog by id",
		)
	}

	return documentToBlog(data), nil
}
