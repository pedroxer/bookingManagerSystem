package pkg

import (
	"context"
	"fmt"
	config "github.com/pedroxer/BookingManagerSystem/internal/configs"
	proto_gen "github.com/pedroxer/BookingManagerSystem/internal/proto_gen/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateResourceClient(config config.ServiceConnectInfo) (proto_gen.ResourceServiceClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	service := proto_gen.NewResourceServiceClient(conn)
	_, err = service.Ping(context.Background(), &proto_gen.PingRequest{})
	if err != nil {
		return nil, err
	}
	return service, nil
}

func CreateBookingClient(config config.ServiceConnectInfo) (proto_gen.BookingServiceClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto_gen.NewBookingServiceClient(conn), nil
}
