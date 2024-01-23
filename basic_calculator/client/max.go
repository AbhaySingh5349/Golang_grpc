package main

import (
	"context"
	"log"
	"time"
	"io"
	"math/rand"

	pb "basic_calculator/compiled_protos/protos" // importing the generated libraries
)

func getMaximum(client pb.CalculatorServiceClient) {
	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("error in opening client stream: %v", err)
	}

	// creating a 'go channel' & will have '2 go routines' running simultaneously
	wait_channel := make(chan struct{})

	// 'go routine' to send all requests that we have in array
	go func() {
		rand.Seed(time.Now().UnixNano())
		size := int(rand.Intn(10))

		var reqs = make([]int, size)
		log.Printf("array sized %d", size)
		for i := 0; i < size; i++ {
			val := int(rand.Intn(100))
			log.Printf("%d ", val)
			reqs[i] = val
		}

		for _, req := range reqs {
			stream.Send(&pb.MaxRequest{Num: int32(req)}) // sending requests to server
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

			log.Printf("Max: %d", res.MaxResponse)
		}
		close(wait_channel)
	}()

	// we need to make channel wait for calling 'close' & let 2 'go routines' send, receive streams
	<-wait_channel
}