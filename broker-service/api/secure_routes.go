package api

import (
	"broker-service/api/middleware"
	"broker-service/internal/grpc/client"
	"net/http"

	"github.com/AdamShannag/toolkit"
	"github.com/go-chi/chi/v5"
)

func SecureRoutes(kit *toolkit.Tools, mid *middleware.Middleware, clients *client.GrpcClients) http.Handler {
	mux := chi.NewRouter()
	mux.Use(mid.VerifyToken)

	// Mount routes here!!

	return mux
}
