package main

import (
	"log"
	"net"
	"os"

	pb "github.com/mualle/go-gRpc-example/proto/todo/v1"
	"google.golang.org/grpc"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}

	addr := args[0]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(lis)

	log.Printf("listening at %s\n", addr)
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	//Link grpc server with our created server
	pb.RegisterTodoServiceServer(s, &server{
		d: New(),
	})

	//registration of endpoints
	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

}
