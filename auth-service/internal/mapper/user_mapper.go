package mapper

import (
	db "auth-service/db/sqlc"
	pb "auth-service/internal/grpc/proto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToResource(user db.User) *pb.User {
	return &pb.User{
		Uuid:       user.Uuid,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		Address:    user.Address,
		UserType:   user.UserType,
		ModifiedAt: timestamppb.New(user.ModifiedAt),
		CreatedAt:  timestamppb.New(user.CreatedAt),
	}
}
