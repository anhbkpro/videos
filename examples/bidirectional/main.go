package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	wearablepb "github.com/anhbkpro/videos/2022/07/23/gen/go/wearable/v1"
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
	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.CalculateBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalln("Opening stream", err)
	}

	for i := 0; i < 10; i++ {
		if err := stream.Send(&wearablepb.CalculateBeatsPerMinuteRequest{
			Value: uint32(i * 2),
		}); err != nil {
			log.Fatalln("Send", err)
		}
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}

		fmt.Println("Average", res.GetAverage())
	}
}
