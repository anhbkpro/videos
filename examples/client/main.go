package main

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/anhbkpro/videos/2022/07/23/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "1-2-3",
	})
	if err != nil {
		log.Fatalf("failed to GetUser: %v", err)
	}

	fmt.Printf("%+v\n", res)
}
