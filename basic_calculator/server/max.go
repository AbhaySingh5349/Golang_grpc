package main

import (
	"log"
	"io"
	"math"

	pb "basic_calculator/compiled_protos/protos"
)

func (s *CalculatorServiceServer) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function is invoked")

	var cur_max int32 = math.MinInt32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// communication should be closed b/w server & client (as client will not send more requests)
			return nil
		}

		if err != nil {
			log.Fatalf("error in client stream reading: %v", err)
		}

		cur_num := req.Num
		if cur_num > cur_max {
			cur_max = cur_num
		}

		err = stream.Send(&pb.MaxResponse{MaxResponse: cur_max})
		if err != nil {
			log.Fatalf("error in sending max to client: %v", err)
		}
	}
}
