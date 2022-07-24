package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	stream, err := client.ConsumeBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalln("Consuming stream", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{
			Uuid:   "anhlai",
			Value:  uint32(i),
			Minute: uint32(i * 2),
		})
		if err != nil {
			log.Fatalln("sending value", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("closing", err)
	}

	fmt.Printf("Total message: %d", res.GetTotal())
}
