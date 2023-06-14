package main

import (
	db "auth-service/db/sqlc"
	"auth-service/internal/config"
	pb "auth-service/internal/grpc/proto/auth"
	"auth-service/internal/grpc/server"
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"net"
)

const webPort = ":8091"

func main() {

	addr := "localhost:50001"
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	newServer, _ := server.NewServer(conf, db.NewStore(conn))
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterAuthServiceServer(s, newServer)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
