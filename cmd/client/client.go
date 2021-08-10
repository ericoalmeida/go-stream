package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	AddUsers(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "001",
		Name:  "Customer",
		Email: "customer@enterprise@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), request)

	if err != nil {
		log.Fatalf("Cound not make gRPC request: %v", err)
	}

	// loop infinito
	for {
		stream, error := responseStream.Recv()

		if error == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Cound not receive the message: %v", err)
		}

		fmt.Println("Status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	requests := []*pb.User{
		{
			Id:    "001",
			Name:  "User 001",
			Email: "user@email.com",
		},
		{
			Id:    "002",
			Name:  "User 002",
			Email: "user@email.com",
		},
		{
			Id:    "003",
			Name:  "User 003",
			Email: "user@email.com",
		},
		{
			Id:    "004",
			Name:  "User 004",
			Email: "user@email.com",
		},
		{
			Id:    "005",
			Name:  "User 005",
			Email: "user@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, request := range requests {
		stream.Send(request)
		time.Sleep(time.Second * 3)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	fmt.Println(response)
}
