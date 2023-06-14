package server

import (
	db "auth-service/db/sqlc"
	"auth-service/internal/security"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) getUser(ctx context.Context, uuid string) (db.User, error) {
	user, err := server.store.GetUser(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.User{}, status.Errorf(codes.NotFound, "invalid credential")
		}
		return db.User{}, status.Errorf(codes.Internal, "failed to find user")
	}
	return user, nil
}

func (server *Server) checkPassword(password string, user db.User) error {
	err := security.CheckPassword(password, user.HashedPassword)
	if err != nil {
		return status.Errorf(codes.NotFound, "invalid credential")
	}
	return nil
}
