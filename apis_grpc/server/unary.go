package main

import (
	"context"
	"fmt"
	pb "apis_grpc/compiled_protos/protos"
)

// we want to implement this function under GreetServiceServer type
func (s *GreetServiceServer) Greet(ctx context.Context,in *pb.GreetRequest) (*pb.GreetResponse, error){
	fmt.Println("GREET FUNC INVOKED")

	fname := in.GetFirstName()

	result := "Hello " + fname

	return &pb.GreetResponse{
		Result: result,
	},nil
}