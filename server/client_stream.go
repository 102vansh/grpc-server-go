package main

import (
	"io"
	"log"

	pb "github.com/grpc/proto"
)

func (s *helloserver) Sayhelloclientstreaming(stream pb.GreetsService_SayhelloclientstreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Messagelist{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("got req with name%v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
