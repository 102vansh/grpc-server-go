package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpc/proto"
)

func callsayhelloclientstream(client pb.GreetsServiceClient, names *pb.Namelist) {
	log.Printf("client streaming strates")
	stream, err := client.Sayhelloclientstreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names%v", err)

	}
	for _, name := range names.Names {
		req := &pb.Hellorequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("sent req with name %s", name)
		time.Sleep(2 * time.Second)

	}
	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil {
		panic(err.Error())
	}
	log.Printf("%v", res.Messages)
}
