package main

import (
	"log"
	"io"
	pb "basic_calculator/compiled_protos/protos"
)

func (s *CalculatorServiceServer) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Average function is invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// communication should be closed b/w server & client (as client will not send more requests)
			return stream.SendAndClose(&pb.AverageResponse{AverageResponse: float64(sum) / float64(count)})
		}

		if err != nil {
			log.Fatalf("error in client stream reading: %v", err)
		}

		sum += req.Num
		count++
	}
}