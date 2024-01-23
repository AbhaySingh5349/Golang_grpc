package main

import (
	"log"
	"context"
	"fmt"
	"math"
	pb "basic_calculator/compiled_protos/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *CalculatorServiceServer) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function is invoked: %v", in)

	val := float64(in.Num)

	if val < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("Received a -ve number: %f", val),
		)
	}

	return &pb.SqrtResponse{SqrtResponse: float32(math.Sqrt(val))}, nil
}