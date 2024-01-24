package main

import (
	"context"
	"fmt"
	"log"

	pb "blog/compiled_protos/protos"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// it has same properties as 'blog objects' in '.proto'
// bson used by mongo db to store doc
type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // ID will be serialized to _id (omit if its empty)
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogID, error) {
	log.Printf("CreateBlog was invoked with %v\n", in)

	// creating 'blog item' instance
	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	// inserting 'blog item' instance to DB & mongoDB will automatcally return unique ID of type 'interface'
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal error: %v", err),
		)
	}

	// casting 'interface ID  --> Object ID'
	oid, isConverted := res.InsertedID.(primitive.ObjectID)

	if !isConverted {
		return nil, status.Errorf(
			codes.Internal, "Cannot convert to Object ID",
		)
	}

	return &pb.BlogID{
		Id: oid.Hex(),
	}, nil
}