package main

import (
	"log"
	"net"

	"github.com/ericoalmeida/grpc/pb"
	"github.com/ericoalmeida/grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Create a listener at 50051 local port
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	//registering services
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	reflection.Register(grpcServer)

	// Initiating a gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Coud not serve: %v", err)
	}
}
