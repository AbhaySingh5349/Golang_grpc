package main

import (
	pb "blog/compiled_protos/protos"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// address of grpc server whcih will be hosted at 'localhost'
	address = "localhost:50051"
)


func main() {
	// dial a connection to grpc
	// WithBlock() means, this connection will not be returned until its made (acting as a blocking call to Dial)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	// id := createNewBlog(client) // returns ID of blog
	// id := `65b0790c9722082c9501b16a` 
	// readNewBlog(client, id) // to pass hardcoded id: `{oid_value}`
	// id := `65b0790c9722082c9501b16a` 
	// updateExistingBlog(client, id)
	// deleteBlog(client, id)
	getAllBlogs(client)
}