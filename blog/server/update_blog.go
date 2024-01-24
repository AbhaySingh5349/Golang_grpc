package main

import (
	"context"
	"log"

	pb "blog/compiled_protos/protos"

	"google.golang.org/protobuf/types/known/emptypb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogServiceServer) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id) // transform ID of blog to Object ID
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, "Cannot parse ID",
		)
	}

	// creating Blog Item instance
	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	// update in db
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal, "Could not update",
		)
	}

	// error related to filter
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound, "Cannot find blog with ID",
		)
	}

	return &emptypb.Empty{}, nil
}