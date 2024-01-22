package main

import (
	"io"
	"log"
	pb "apis_grpc/compiled_protos/protos"

)

// we want to implement this function under GreetServiceServer type
func (s *GreetServiceServer) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// communication should be closed b/w server & client (as client will not send more requests)
			return nil
		}

		if err != nil {
			log.Fatalf("error in client stream reading: %v", err)
		}

		result := "hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{Result: result})

		if err != nil {
			log.Fatalf("error while sending data to client: %v", err)
		}
	}
}