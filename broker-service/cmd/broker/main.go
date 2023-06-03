package main

import (
	"broker-service/api"
	"broker-service/internal/shutdown"
	"context"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	var (
		ctx    = context.Background()
		mux    = api.NewMux()
		server = http.Server{
			Addr:    ":" + webPort,
			Handler: mux,
		}
		shutdownChan = make(chan struct{})
	)

	go shutdown.GracefulShutdown(ctx, &server, shutdownChan)

	log.Println("server starting: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("server error", err)
	}

	<-shutdownChan
}
