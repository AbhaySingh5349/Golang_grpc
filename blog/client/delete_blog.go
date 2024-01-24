package main

import (
	"log"
	"context"
	
	pb "blog/compiled_protos/protos"
)

func deleteBlog(client pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")
	_, err := client.DeleteBlog(context.Background(), &pb.BlogID{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}