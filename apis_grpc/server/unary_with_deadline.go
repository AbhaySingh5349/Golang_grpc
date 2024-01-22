package main

import (
	"context"
	"log"
	"time"
	pb "apis_grpc/compiled_protos/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// we want to implement this function under GreetServiceServer type
func (s *GreetServiceServer) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline function was invoked: %v", in)

	// loop to check if deadline was exceeded (each loop we wait for 1 sec & here total wait time be 3 sec)
	// if deadline was not exceeded after 3 sec, we return a 'greet response'

	for t := 0; t < 3; t++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Client cancelled request as deadline was exceeded")
			return nil, status.Error(codes.Canceled, "Client cancelled request as deadline was exceeded")
		}
		time.Sleep(1 * time.Second)
	}

	fname := in.GetFirstName()
	result := "hello " + fname

	return &pb.GreetResponse{
		Result: result,
	}, nil
}