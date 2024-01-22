package main

import (
	"context"
	"log"
	"io"
	pb "apis_grpc/compiled_protos/protos"

)

// we want to implement this function under GreetServiceServer type
func responseToGreetManyTimes(client pb.GreetServiceClient) {
	// calling 'rpc endpoint' (we will receive a response or an error)
	req := &pb.GreetRequest{FirstName: "Engineer Bassi server streaming"}
	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error in greeting many times: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			// communication should be closed b/w server & client
			break
		}
		if err != nil {
			log.Fatalf("error in stream reading: %v", err)
		}

		log.Printf("Greeting: %s", res.Result)
	}
}