package main

import (
	"context"
	"log"
	"net"
	"time"

	userpb "github.com/anhbkpro/videos/2022/07/23/gen/go/user/v1"
	wearablepb "github.com/anhbkpro/videos/2022/07/23/gen/go/wearable/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type userService struct {
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     req.Uuid,
			FullName: "Anh",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServer := &userService{}
	wearableServer := &wearableService{}

	healthServer := health.NewServer()

	go func() {
		for {
			status := healthpb.HealthCheckResponse_SERVING
			// check if UserService is valid
			if time.Now().Second()%2 == 0 {
				status = healthpb.HealthCheckResponse_NOT_SERVING
			}

			healthServer.SetServingStatus(userpb.UserService_ServiceDesc.ServiceName, status)

			time.Sleep(1 * time.Second)
		}
	}()

	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(userpb.UserService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	userpb.RegisterUserServiceServer(grpcServer, userServer)
	wearablepb.RegisterWearableServiceServer(grpcServer, wearableServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	// we should not enable this on production
	// reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
