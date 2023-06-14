package server

import (
	db "auth-service/db/sqlc"
	pb "auth-service/internal/grpc/proto/auth"
	"auth-service/internal/mapper"
	"auth-service/internal/security"
	val "auth-service/internal/validator"
	"context"
	"database/sql"
	"errors"
	"golang.org/x/exp/slices"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if err := checkUserAndPassword(ctx, req, server); err != nil {
		return nil, err
	}

	arg := db.UpdateUserParams{
		Uuid: req.GetUuid(),
		Username: sql.NullString{
			String: req.GetUsername(),
			Valid:  req.Username != "",
		},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  req.Phone != "",
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != "",
		},
		Address: sql.NullString{
			String: req.GetAddress(),
			Valid:  req.Address != "",
		},
		ModifiedAt: server.time.Now(),
	}

	if slices.Contains(req.Payload.Roles, "ADMIN") {
		arg.UserType = sql.NullString{
			String: req.GetUserType(),
			Valid:  req.UserType != "",
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UserResponse{
		User: mapper.UserToResource(user),
	}
	return rsp, nil
}

func checkUserAndPassword(ctx context.Context, req *pb.UserRequest, server *Server) error {
	user, err := server.store.GetUser(ctx, req.GetUuid())
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "user not found")
		}
		return status.Errorf(codes.Internal, "failed to update user: %s", err)
	}
	if user.Uuid == req.Payload.GetUuid() {
		return checkPassword(err, req, user)
	}
	return nil
}

func checkPassword(err error, req *pb.UserRequest, user db.User) error {
	err = security.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return status.Errorf(codes.NotFound, "incorrect password")
	}
	return nil
}

func validateUpdateUserRequest(req *pb.UserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Uuid == "" {
		violations = append(violations, fieldViolation("uuid", errors.New("can't be empty")))
	}
	if req.Payload == nil {
		violations = append(violations, fieldViolation("payload", errors.New("can't be empty")))
	}

	if req.Username != "" {
		if err := val.ValidateUsername(req.GetUsername()); err != nil {
			violations = append(violations, fieldViolation("username", err))
		}
	}

	if req.Email != "" {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	return violations
}
