package main

import (
	"math/rand"
	"log"
	"time"
	"fmt"
	pb "apis_grpc/compiled_protos/protos"

)

// we want to implement this function under GreetServiceServer type

func (s *GreetServiceServer) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked: %v", in)

	rand.Seed(time.Now().UnixNano())
	count := int(rand.Intn(10))

	for i := 0; i < count; i++ {
		result := fmt.Sprintf("hello %s, count %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: result,
		})
	}
	return nil
}