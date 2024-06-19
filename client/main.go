package main

import (
	"fmt"
	"log"

	pb "github.com/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8085"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetsServiceClient(conn)
	fmt.Println(client)

	names := &pb.Namelist{
		Names: []string{"akhil", "Alice", "bob"},
	}

	Callsayhelloserverstream(client, names)
	callsayhelloclientstream(client, names)
	Callsayhello(client)

}
