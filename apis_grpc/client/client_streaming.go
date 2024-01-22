package main

import (
	"log"
	"context"
	"time"
	pb "apis_grpc/compiled_protos/protos"

)

// we want to implement this function under GreetServiceServer type
func responseToLongGreet(client pb.GreetServiceClient) {
	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while creating LongGreet stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Aakash"},
		{FirstName: "Aayush"},
		{FirstName: "Engineer bassi"},
	}

	for _, req := range reqs {
		stream.Send(req) // sending requests to server
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}