package main

import (
	"io"
	"log"
	"time"
	"context"
	pb "apis_grpc/compiled_protos/protos"

)

func responseToGreetEveryone(client pb.GreetServiceClient) {
	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating GreetEveryone stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Aakash"},
		{FirstName: "Aayush"},
		{FirstName: "Abhay"},
	}

	// creating a 'go channel' & will have '2 go routines' running simultaneously
	wait_channel := make(chan struct{})

	// 'go routine' to send all requests that we have in array
	go func() {
		for _, req := range reqs {
			stream.Send(req) // sending requests to server
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// 'go routine' will receive responses from server
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				// communication should be closed b/w server & client
				break
			}
			if err != nil {
				log.Fatalf("error in server stream reading: %v", err)
				break
			}

			log.Printf("Greeting: %s", res.Result)
		}
		close(wait_channel)
	}()

	// we need to make channel wait for calling 'close' & let 2 'go routines' send, receive streams
	<-wait_channel
}