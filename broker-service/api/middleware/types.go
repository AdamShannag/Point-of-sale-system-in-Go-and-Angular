package middleware

import (
	"broker-service/internal/grpc/proto/auth"

	"github.com/AdamShannag/toolkit"
)

type Middleware struct {
	kit        *toolkit.Tools
	authClient auth.AuthServiceClient
}

func NewMiddleware(kit *toolkit.Tools, authClient auth.AuthServiceClient) *Middleware {
	return &Middleware{kit, authClient}
}
