package main

import (
	"log"
	"io"
	"fmt"
	pb "apis_grpc/compiled_protos/protos"

)

// we want to implement this function under GreetServiceServer type
func (s *GreetServiceServer) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	result := "" // concatenate all requests

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// communication should be closed b/w server & client (as client will not send more requests)
			return stream.SendAndClose(&pb.GreetResponse{Result: result})
		}

		if err != nil {
			log.Fatalf("error in client stream reading: %v", err)
		}

		result += fmt.Sprintf("hello %s!\n", req.FirstName)
	}
}