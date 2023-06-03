package api

import (
	"broker-service/api/handler/auth"
	"broker-service/api/middleware"
	"broker-service/internal/grpc/client"

	"github.com/AdamShannag/toolkit"
	chimid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewMux() *chi.Mux {
	var (
		mux     = chi.NewMux()
		kit     = &toolkit.Tools{}
		clients = client.New()
		mid     = middleware.NewMiddleware(kit, *clients.AuthClient)
	)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(chimid.Heartbeat("/ping"))
	mux.Use(chimid.RequestID)
	mux.Use(chimid.RealIP)
	mux.Use(chimid.Recoverer)

	mux.Mount("/api/auth", auth.NewAuth(kit))
	mux.Mount("/api", SecureRoutes(kit, mid, clients))

	return mux
}
