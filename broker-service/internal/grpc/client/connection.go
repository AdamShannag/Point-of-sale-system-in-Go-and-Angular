package client

import (
	"broker-service/internal/shutdown"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	shutdown.AddToShutdowns(func() error {
		conn.Close()
		return nil
	})

	return conn
}
