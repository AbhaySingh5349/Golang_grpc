package main

import (
	"context"
	"log"
	"time"
	pb "apis_grpc/compiled_protos/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func responseToGreetWithDeadline(client pb.GreetServiceClient, t time.Duration) {

	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout*int32(time.Second)))
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel() // cancel will be called if deadline exceeded

	req := &pb.GreetRequest{FirstName: "Engineer Bassi greet witin 10 seconds"}
	res, err := client.GreetWithDeadline(ctx, req)

	if err != nil {
		// transforming error to 'grpc status'
		e, isGrpcError := status.FromError(err)

		if isGrpcError {
			log.Printf("Error message from server: %s", e.Message())
			log.Printf("Error code from server: %s", e.Code())

			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded")
				return
			} else {
				log.Fatalf("unexpected grpc error : %v", err)
			}
		} else {
			log.Fatalf("non-grpc error in greeting with deadline: %v", err)
		}
	}
	log.Printf("Greet with deadline: %s", res.Result)
}