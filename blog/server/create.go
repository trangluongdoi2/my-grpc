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

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog function was envoke %v\n", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	// Debug: log the actual type and value of InsertedID
	log.Printf("InsertedID value: %v, type: %T\n", res.InsertedID, res.InsertedID)

	// Try type assertion
	oid, ok := res.InsertedID.(bson.ObjectID)

	if !ok {
		// Log detailed error with actual type
		log.Printf("Type assertion failed! Expected bson.ObjectID, got %T\n", res.InsertedID)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID, got type: %T, value: %v\n", res.InsertedID, res.InsertedID),
		)
	}

	log.Printf("Successfully converted to ObjectID: %s\n", oid.Hex())

	return &pb.BlogId{Id: oid.Hex()}, nil
}
