package main

import (
	proto "auth-service/internal/grpc/proto/auth"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	authClient := proto.NewAuthServiceClient(newConnection("localhost:50001"))
	//user, err := authClient.CreateUser(context.Background(), &proto.UserRequest{
	//	Username: "adam",
	//	Email:    "adam@gmail.com",
	//	Phone:    "07555",
	//	Address:  "amman",
	//	Password: "P@ssw0rd",
	//	UserType: "USER",
	//	Payload: &proto.Payload{
	//		Roles: []string{},
	//		Uuid:  "",
	//	},
	//})
	//
	//user, err := authClient.Login(context.Background(), &proto.LoginRequest{
	//	Username: "mohammad",
	//	Password: "Moha@555",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//payload, err := authClient.VerifyToken(context.Background(), &proto.VerifyTokenRequest{
	//	Token: user.AccessToken})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(payload)

	user, err := authClient.UpdateUser(context.Background(), &proto.UserRequest{
		Uuid:     "1a9d06e0-6387-4bc5-bffb-ad7ee44de63d",
		Username: "mohammad",
		UserType: "ADMIN",
		Payload: &proto.Payload{
			Username: "mohammad",
			Uuid:     "1a9d06e0-6387-4b c5-bffb-ad7ee44de63d",
			Roles:    []string{"ADMIN"},
		},
		//Password: "Moha@555",
	})

	changed, err := authClient.ChangePassword(context.Background(), &proto.ChangePasswordRequest{
		Username:    "mohammad",
		NewPassword: "Moha@1234",
		//OldPassword: "Moha@1234",
		Payload: &proto.Payload{
			Username: "mohammad",
			Uuid:     "1a9d06e0-6387-4bc5-bff b-ad7ee44de63d",
			Roles:    []string{"ADMIN"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(changed)

	log.Println(user)

}

func newConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
