package client

import "os"

const (
	auth = "AUTH_GRPC_SERVER"
	// Add other grpc addresses
)

func getAddress(addr string) string {
	return os.Getenv(addr)
}
