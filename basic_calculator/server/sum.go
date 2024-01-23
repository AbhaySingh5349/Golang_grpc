package main

import (
	"log"
	"context"
	pb "basic_calculator/compiled_protos/protos"
)

func (s *CalculatorServiceServer) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function is invoked: %v", in)
	a, b := in.GetFirstNum(), in.GetSecondNum()

	result := a + b

	return &pb.SumResponse{SumResponse: result}, nil
}