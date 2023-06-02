package middleware

import "broker-service/internal/grpc/client"

type Middleware struct {
	client *client.GrpcClients
}

func NewMiddleware(authClient *client.GrpcClients) *Middleware {
	return &Middleware{client: authClient}
}
