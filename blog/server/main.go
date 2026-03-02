package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// Connect to MongoDB with authentication (v2 API - no context param)
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v\n", err)
	}

	// Ensure disconnect on exit
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v\n", err)
		}
	}()

	// Verify connection with ping
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v\n", err)
	}
	log.Println("✓ Connected to MongoDB successfully!")

	// Set up collection
	collection = client.Database("blogdb").Collection("blog")
	log.Printf("Using database: blogdb, collection: blog\n")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to server: %v\n", err)
	}
}
