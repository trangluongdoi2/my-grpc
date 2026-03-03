package main

import (
	pb "github.com/trangluongdoi2/my-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type BlogItem struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	AuthorId string        `bson:"author_id"`
	Title    string        `bson:"title"`
	Content  string        `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
