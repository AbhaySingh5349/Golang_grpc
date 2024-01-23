package main

import (
	"context"
	"log"
	"fmt"

	pb "basic_calculator/compiled_protos/protos" // importing the generated libraries

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func getSqrt(client pb.CalculatorServiceClient) {
	var val int32
	fmt.Printf("i/p number to get its sqrt: ")
	fmt.Scan(&val)

	req := &pb.SqrtRequest{Num: val}
	res, err := client.Sqrt(context.Background(), req)

	if err != nil {
		// transforming error to 'grpc status'
		e, isGrpcError := status.FromError(err)

		if isGrpcError {
			log.Printf("Error message from server: %s", e.Message())
			log.Printf("Error code from server: %s", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("-ve number was i/p by client")
				return
			} else {
				log.Fatalf("unexpected grpc error : %v", err)
			}
		} else {
			log.Fatalf("non-grpc error in calculating sqrt: %v", err)
		}
	}
	log.Printf("Sqrt is: %f", res.SqrtResponse)
}