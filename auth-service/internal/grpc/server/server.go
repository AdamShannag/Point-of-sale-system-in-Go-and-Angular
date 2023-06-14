package server

import (
	db "auth-service/db/sqlc"
	"auth-service/internal/config"
	pb "auth-service/internal/grpc/proto/auth"
	"auth-service/internal/log"
	"auth-service/internal/time"
	"auth-service/internal/token"
	"fmt"
	j "github.com/dgrijalva/jwt-go"
	"github.com/mohammadyaseen2/TokenUtils/jwt"
	"github.com/mohammadyaseen2/TokenUtils/model"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	config config.Config
	store  db.Store
	token  token.TokenMaker
	time   time.TimeService
	log    *log.Logger
	pb.AuthServiceServer
}

// NewServer creates a new gRPC server.
func NewServer(config config.Config, store db.Store) (*Server, error) {
	fmt.Println("NewServer")
	keyPair, err := getKeyPair(config)
	if err != nil {
		return nil, err
	}
	tokenMaker := getTokenMaker(keyPair)
	timeService, err := time.NewTimeService(config.AppTimeZone)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
		time:   *timeService,
		log:    log.NewLogger(),
	}

	return server, nil
}

func getTokenMaker(keyPair *model.KeyPair) token.TokenMaker {
	newJwt := jwt.New(keyPair.PrivateKey, *j.SigningMethodRS256)
	validator := jwt.NewValidator(keyPair.PublicKey)

	return token.NewJWTMaker(*newJwt, *validator)
}

func getKeyPair(config config.Config) (*model.KeyPair, error) {
	keysMaker := token.NewKeysMaker(config.RsaKeysBits)
	keyPair, err := keysMaker.GetKeyPair(config.PrivateKeyFilePath, config.PublicKeyFilePath)
	if err != nil {
		return nil, fmt.Errorf("error while get kay pair: %w", err)
	}
	return keyPair, nil
}
