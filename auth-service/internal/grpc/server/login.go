package server

import (
	db "auth-service/db/sqlc"
	proto "auth-service/internal/grpc/proto/auth"
	"auth-service/internal/mapper"
	val "auth-service/internal/validator"
	"context"
	"github.com/mohammadyaseen2/TokenUtils/model"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	if violations := validateLoginUserRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.getUser(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	if err := server.checkPassword(req.GetPassword(), user); err != nil {
		return nil, err
	}

	payload := server.createPayload(user, server.config.AccessTokenDuration)
	accessToken, err := server.token.CreateToken(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	payload = server.createPayload(user, server.config.RefreshTokenDuration)
	refreshToken, err := server.token.CreateToken(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	convertedUser := mapper.UserToResource(user)
	if user.AddedBy.Valid {
		convertedUser.AddedBy = getUserWhoAdded(ctx, server, user.AddedBy.String)
	}

	rsp := &proto.LoginResponse{
		User:         convertedUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return rsp, nil
}

func (server *Server) createPayload(user db.User, duration time.Duration) *model.Payload {
	payload, _ := model.NewPayload(
		user.Username,
		[]string{user.UserType},
		duration,
		model.Extra{"uuid": user.Uuid},
	)
	payload.IssuedAt = server.time.Now().Unix()
	payload.ExpiresAt = server.time.Now().Add(duration).Unix()
	return payload
}

func getUserWhoAdded(ctx context.Context, server *Server, userUuid string) string {
	user, _ := server.store.GetUsername(ctx, userUuid)
	return user
}

func validateLoginUserRequest(req *proto.LoginRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
