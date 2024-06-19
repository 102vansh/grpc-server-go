package main

import (
	"log"
	"net"

	pb "github.com/grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8085"
)

type helloserver struct {
	pb.GreetsServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failes to start the server%v", err)
	}

	grpcserver := grpc.NewServer()
	pb.RegisterGreetsServiceServer(grpcserver, &helloserver{})
	log.Printf("server started t %v", lis.Addr())
	if err := grpcserver.Serve(lis); err != nil {
		log.Fatal("failed to start")
	}
}
