package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/trangluongdoi2/my-grpc/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to server: %v\n", err)
	}
}
