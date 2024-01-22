package main

import (
	"context"
	"fmt"
	"log"
	pb "apis_grpc/compiled_protos/protos"

	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

type GreetServiceServer struct{
	pb.UnimplementedGreetServiceServer
}

func (s *GreetServiceServer) Greet(ctx context.Context,in *pb.GreetRequest) (*pb.GreetResponse, error){
	fmt.Println("GREET FUNC IS INVOKED after makefile is success with clean make & final")

	fname := in.GetFirstName()

	result := "Hello " + fname

	return &pb.GreetResponse{
		Result: result,
	},nil
}

func main(){
	listener, err := net.Listen("tcp", port) // function to begin listening the port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                                   // creating 'server object' from grpc module
	pb.RegisterGreetServiceServer(s, &GreetServiceServer{}) // register server as new grpc service & we also pass reference to server structure type
	reflection.Register(s)                                  // for serializing & de-serializing data

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to load server: %v", err)
	}
}