package main

import (
	"context"
	"log"
	"net"

	pb "blog/compiled_protos/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// port on which grpc server will run on
	port = ":50051"
)

type BlogServiceServer struct {
	pb.UnimplementedBlogServiceServer
}

// client will connect to db which will be used to access collection
var collection *mongo.Collection

func main() {
	// client will connect to mongoDB database & we can access collection 'blog'
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/")) // docker is binding mongoDB on system at port 27017
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background()) // mongo connection is an Object that will help in CRUD operation
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog") // name of DB & Collection (creating or retrieving db & collection)

	listener, err := net.Listen("tcp", port) // function to begin listening the port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // creating 'server object' from grpc module
	pb.RegisterBlogServiceServer(s, &BlogServiceServer{}) // register server as new grpc service & we also pass reference to server structure type
	reflection.Register(s) // for server to expose which endpoints are available & allow CLI to talk to server without priliminary .proto file

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to load server: %v", err)
	}
}