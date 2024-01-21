
package main

import (
	"context"
	"log"

	pb "apis_grpc/compiled_protos/proto" // importing the generated librarys

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// address of grpc server whcih will be hosted at 'localhost'
	address = "localhost:8081"
)

func responseToGreet(client pb.GreetServiceClient) {
	// calling 'rpc endpoint' (we will receive a response or an error)
	req := &pb.GreetRequest{FirstName: "Abhay Singh"}
	res, err := client.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error in greeting: %v", err)
	}
	log.Printf("Greeting: %s", res.Result)
}

func main() {
	// dial a connection to grpc
	// WithBlock() means, this connection will not be returned until its made (acting as a blocking call to Dial)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn) // client is returned that is used to call 'rpc endpoint'

	responseToGreet(client)
}