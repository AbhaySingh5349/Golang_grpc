package main

import (
	"context"
	"log"
	"time"
	"io"
	"math"
	"math/rand"

	pb "basic_calculator/compiled_protos/protos" // importing the generated libraries
)


func getPrimes(client pb.CalculatorServiceClient) {
	// calling 'rpc endpoint' (we will receive a response or an error)

	rand.Seed(time.Now().UnixNano())
	num_range := uint32(rand.Intn(int(math.Pow(2, 32) - 1)))

	log.Printf("Primes over range: %d", num_range)

	req := &pb.PrimesRequest{Range: num_range}
	stream, err := client.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error in generating primes: %v", err)
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

		log.Printf("Primes: %d", res.PrimesResponse)
	}
}