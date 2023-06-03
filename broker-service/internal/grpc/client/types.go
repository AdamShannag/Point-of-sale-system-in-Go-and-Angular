package client

import (
	pb "broker-service/internal/grpc/proto/auth"
)

type GrpcClients struct {
	AuthClient *pb.AuthServiceClient
}

func New() *GrpcClients {
	authClient := pb.NewAuthServiceClient(newConnection(getAddress(auth)))
	return &GrpcClients{
		AuthClient: &authClient,
	}
}
