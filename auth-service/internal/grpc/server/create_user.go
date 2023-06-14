package server

import (
	db "auth-service/db/sqlc"
	proto "auth-service/internal/grpc/proto/auth"
	"auth-service/internal/mapper"
	"auth-service/internal/security"
	val "auth-service/internal/validator"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := security.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}
	fmt.Println("req.Payload.Uuid: ", req.Payload.Uuid)

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Uuid:           uuid.New().String(),
			Username:       req.GetUsername(),
			Email:          req.GetEmail(),
			Phone:          req.GetPhone(),
			UserType:       req.GetUserType(),
			HashedPassword: hashedPassword,
			Address:        req.GetAddress(),
			AddedBy: sql.NullString{
				Valid:  req.Payload.Uuid != "",
				String: req.Payload.GetUuid(),
			},
			CreatedAt:  server.time.Now(),
			ModifiedAt: server.time.Now(),
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &proto.UserResponse{
		User: mapper.UserToResource(txResult.User),
	}
	return rsp, nil
}

func validateCreateUserRequest(req *proto.UserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}
