package main

import (
	"context"

	pb "github.com/grpc/proto"
)

func (s *helloserver) sayhello(ctx context.Context, req *pb.Noparams) (*pb.Helloresponse, error) {
	return &pb.Helloresponse{
		Message: "Hello",
	}, nil
}
