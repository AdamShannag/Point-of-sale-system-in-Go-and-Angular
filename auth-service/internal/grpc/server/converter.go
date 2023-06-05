package server

import (
	db "auth-service/db/sqlc"
	proto "auth-service/internal/grpc/proto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//string addedBy = 8

func convertUser(user db.User) *proto.User {
	return &proto.User{
		Uuid:       user.Uuid,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		Address:    user.Address,
		UserType:   user.UserType,
		AddedBy:    user.AddedBy,
		ModifiedAt: timestamppb.New(user.ModifiedAt),
		CreatedAt:  timestamppb.New(user.CreatedAt),
	}
}
