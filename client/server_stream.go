package main

import (
	"context"
	"io"
	"log"

	pb "github.com/grpc/proto"
)

func Callsayhelloserverstream(client pb.GreetsServiceClient, names *pb.Namelist) {
	log.Printf("streaming started")
	stream, err := client.Sayhelloserverstreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming", err)

		}

		log.Println(message)
	}
	log.Printf("streaming finished")
}
