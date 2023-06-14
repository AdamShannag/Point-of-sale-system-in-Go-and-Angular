package server

import (
	db "auth-service/db/sqlc"
	pb "auth-service/internal/grpc/proto/auth"
	"auth-service/internal/security"
	val "auth-service/internal/validator"
	"context"
	"errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {

	if violations := validateUpdatePasswordRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.getUser(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	if req.Payload.GetUuid() == user.Uuid {
		if err := server.checkPassword(req.GetOldPassword(), user); err != nil {
			return nil, err
		}
	}

	password, err := security.HashPassword(req.NewPassword)
	rows, err := server.store.UpdatePassword(ctx, db.UpdatePasswordParams{
		Username:       req.Username,
		HashedPassword: password,
		ModifiedAt:     server.time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update password for user [%s]: %s", req.GetUsername(), err)
	}
	return &pb.ChangePasswordResponse{
		Status: rows == 1,
	}, nil
}

func validateUpdatePasswordRequest(req *pb.ChangePasswordRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Payload == nil {
		violations = append(violations, fieldViolation("payload", errors.New("can't be empty")))
	}

	if err := val.ValidatePassword(req.GetNewPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}
