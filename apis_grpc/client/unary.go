package main

import (
	"context"
	"log"

	pb "apis_grpc/compiled_protos/protos" // importing the generated libraries
)

// calls greet rpc endpoint
func responseToUnary(client pb.GreetServiceClient) {
	// calling 'rpc endpoint' (we will receive a response or an error)
	req := &pb.GreetRequest{FirstName: "Engineer Bassi, I'm SDE from IIT(ISM) Dhanbad"}
	res, err := client.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error in greeting: %v", err)
	}
	log.Printf("Greeting: %s", res.Result)
}