package main

import (
	"context"
	"log"
	"fmt"

	pb "blog/compiled_protos/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// mapping function to map 'blog item' to 'Blog object' defined in '.proto' file
func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(), // Hex() helps in converting to 'string'
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func (s *BlogServiceServer) ReadBlog(ctx context.Context, in *pb.BlogID) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", in)

	// get blog id
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, "Cannot parse ID",
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return documentToBlog(data), nil
}