package main

import (
	"broker-service/api"
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
		shutdown = make(chan struct{})
	)

	go gracefulShutdown(ctx, &server, shutdown)

	log.Println("server starting: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("server error", err)
	}

	<-shutdown
}
