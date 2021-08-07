package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ericoalmeida/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "001",
		Name:  "Customer",
		Email: "customer@enterprise@email.com",
	}

	response, err := client.AddUser(context.Background(), request)

	if err != nil {
		log.Fatalf("Cound not make gRPC request: %v", err)
	}

	fmt.Println(response)
}
