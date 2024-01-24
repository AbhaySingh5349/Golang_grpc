package main

import (
	"context"
	"log"
	"fmt"

	pb "blog/compiled_protos/protos"

	"google.golang.org/protobuf/types/known/emptypb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogServiceServer) BlogList(_ *emptypb.Empty, stream pb.BlogService_BlogListServer) error {
	log.Println("ListBlogs was invoked")

	ctx := context.Background()
	cursor, err := collection.Find(ctx, primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal, fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cursor.Close(ctx) // at the end of function, cursor gets closed

	// iterating over cursor
	for cursor.Next(ctx) {
		data := &BlogItem{}
		err := cursor.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal, fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}