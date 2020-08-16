package main

import (
	"log"
	"net"

	"github.com/xiashura/server-jwt-example/pkg/api"
	"google.golang.org/grpc"

	"github.com/xiashura/server-jwt-example/pkg/auth"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := &auth.Client{}
	s := grpc.NewServer()
	api.RegisterAuthServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
