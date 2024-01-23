package main

import (
	"log"
	"context"
	"math/rand"
	"time"

	pb "basic_calculator/compiled_protos/protos"
)

func getAverage(client pb.CalculatorServiceClient) {
	// calling 'rpc endpoint' (we will receive a response or an error)

	stream, err := client.Avg(context.Background())

	if err != nil {
		log.Fatalf("error in opening client stream: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	size := int(rand.Intn(10))

	var arr = make([]int, size)
	log.Printf("array sized %d", size)
	for i := 0; i < size; i++ {
		val := int(rand.Intn(100))
		log.Printf("%d ", val)
		arr[i] = val
	}

	for _, val := range arr {
		stream.Send(&pb.AverageReguest{
			Num: int32(val),
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from AverageReguest: %v", err)
	}

	log.Printf("Average is: %f", res.AverageResponse)
}