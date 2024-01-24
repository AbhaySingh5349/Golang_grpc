package main

import (
	"context"
	"log"

	pb "blog/compiled_protos/protos"
)

// creating blog -> sending to server -> server returns blog id or error
func createNewBlog(client pb.BlogServiceClient) string {
	log.Println("---createNewBlog was invoked---")

	// instance of blog
	blog := &pb.Blog{
		AuthorId: "Abhay",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}