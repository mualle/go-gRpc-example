package main

import (
	pb "github.com/mualle/go-gRpc-example/proto/todo/v1"
)

type server struct {
	d db
	pb.UnimplementedTodoServiceServer
}
