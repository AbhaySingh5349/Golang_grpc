package main

import (
	"log"
	"context"
	
	pb "blog/compiled_protos/protos"
)

func readNewBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogID{Id: id}
	res, err := client.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}