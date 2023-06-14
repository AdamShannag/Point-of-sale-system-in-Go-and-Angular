package server

import (
	pb "auth-service/internal/grpc/proto/auth"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	payload, err := server.token.VerifyToken(req.GetToken())

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "invalid token: %s", err)
	}

	return &pb.VerifyTokenResponse{
		Payload: &pb.Payload{
			Uuid:     fmt.Sprintf("%s", payload.Extra["uuid"]),
			Username: payload.Username,
			Roles:    payload.Roles,
		},
	}, nil
}
