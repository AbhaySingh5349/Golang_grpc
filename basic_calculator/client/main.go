
package main

import (
	"log"

	pb "basic_calculator/compiled_protos/protos" // importing the generated libraries

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// "time"
)

const (
	// address of grpc server whcih will be hosted at 'localhost'
	address = "localhost:8081"
)


func main() {
	// dial a connection to grpc
	// WithBlock() means, this connection will not be returned until its made (acting as a blocking call to Dial)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	// getSum(client)
	// getSqrt(client)
		// getPrimes(client)
	// getAverage(client)
	getMaximum(client)
}