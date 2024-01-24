package main

import (
	"context"
	"log"
	"fmt"

	pb "blog/compiled_protos/protos"

	"google.golang.org/protobuf/types/known/emptypb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *BlogServiceServer) DeleteBlog(ctx context.Context, in *pb.BlogID) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id) // transform ID of blog to Object ID
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, "Cannot parse ID",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	// fliter error
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}