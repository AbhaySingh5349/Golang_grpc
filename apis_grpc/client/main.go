
package main

import (
	"log"

	pb "apis_grpc/compiled_protos/protos" // importing the generated libraries

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"time"
)

const (
	// address of grpc server whcih will be hosted at 'localhost'
	address = "localhost:8081"
)

func main() {
	// dial a connection to grpc
	// since grpc is secure, so it uses SSL by default (but since we dont have credentials to connect with server, so we skip SSL here)
	// WithBlock() means, this connection will not be returned until its made (acting as a blocking call to Dial)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // executes in end to close connection

	client := pb.NewGreetServiceClient(conn) // client is returned that is used to call 'rpc endpoint'

	// responseToUnary(client)
	responseToGreetWithDeadline(client, 10*time.Second)
}