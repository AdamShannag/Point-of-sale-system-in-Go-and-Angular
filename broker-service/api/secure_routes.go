package api

import (
	"broker-service/api/middleware"
	"net/http"

	"github.com/AdamShannag/toolkit"
	"github.com/go-chi/chi/v5"
)

func SecureRoutes(kit *toolkit.Tools) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.VerifyToken)

	// Mount routes here!!

	return mux
}
