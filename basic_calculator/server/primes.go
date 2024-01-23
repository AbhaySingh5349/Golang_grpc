package main

import (
	"log"
	pb "basic_calculator/compiled_protos/protos"
)


func (s *CalculatorServiceServer) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function is invoked: %v", in)

	numRange := in.GetRange()
	divisor := uint32(2)

	for numRange > 1 {
		c := 0
		for numRange%divisor == 0 {
			c++
			numRange /= divisor
		}
		if c > 0 {
			stream.Send(&pb.PrimesResponse{
				PrimesResponse: divisor,
			})
		}
		divisor++
	}
	return nil
}