package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "basic_calculator/compiled_protos/protos" // importing the generated libraries
)

func getSum(client pb.CalculatorServiceClient) {
	// calling 'rpc endpoint' (we will receive a response or an error)

	rand.Seed(time.Now().UnixNano())
	a := int32(rand.Intn(100))
	b := int32(rand.Intn(100))

	log.Printf("a = %v", a)
	log.Printf("b = %v", b)

	req := &pb.SumRequest{FirstNum: int32(a), SecondNum: int32(b)}
	res, err := client.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error in calculating sum: %v", err)
	}
	log.Printf("Sum is: %d", res.SumResponse)
}