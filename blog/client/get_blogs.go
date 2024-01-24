package main

import (
	"context"
	"log"
	"io"

	pb "blog/compiled_protos/protos"

	"google.golang.org/protobuf/types/known/emptypb"
)

func getAllBlogs(client pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")
	stream, err := client.BlogList(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}