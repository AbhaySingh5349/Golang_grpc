package main

import (
	"context"
	"log"

	pb "blog/compiled_protos/protos"
)

func updateExistingBlog(client pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Changed Abhay Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	// empty response
	_, err := client.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}