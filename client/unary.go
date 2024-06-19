package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpc/proto"
)

func Callsayhello(client pb.GreetsServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Sayhello(ctx, &pb.Noparams{})
	if err != nil {
		log.Fatalf("could not greet", err)
	}
	log.Printf("%s", &res.Message)
}
