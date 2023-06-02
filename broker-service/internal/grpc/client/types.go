package client

type GrpcClients struct {
	// AuthClient *AuthServiceClient
}

func New() *GrpcClients {
	// authConn := pb.NewAuthServiceClient(newConnection(AUTH_ADDRESS))

	return &GrpcClients{}
}
