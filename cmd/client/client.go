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
	AddUserStreamBoth(client)
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

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

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

	wait := make(chan int)

	// The go anonymous functions will work in parallel

	go func() {
		// This is a "Thread" that will send data to server
		for _, request := range requests {
			fmt.Println("Sending user:", request.Name)

			stream.Send(request)

			time.Sleep(time.Second * 3)
		}

		stream.CloseSend()
	}()

	go func() {
		// This is a "Thread" that will receive data from server
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving stream from the server: %v", err)
			}

			fmt.Printf("Recebendo user %v com status: %v\n", response.GetUser().GetName(), response.GetStatus())
		}

		close(wait)
	}()

	<-wait
}
