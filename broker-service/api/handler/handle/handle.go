package handle

import (
	"net/http"

	"github.com/AdamShannag/toolkit"
	"github.com/go-chi/chi/v5"
)

type Handle struct {
	*chi.Mux
	kit *toolkit.Tools
}

func NewHandle(toolkit *toolkit.Tools) Handle {
	h := Handle{
		Mux: chi.NewMux(),
		kit: toolkit,
	}

	h.HandleFunc("/{service}", h.HandleRoutes)

	return h
}

func (h *Handle) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	vals := chi.URLParam(r, "service")
	if err := h.kit.WriteJSON(w, 200, toolkit.JSONResponse{Error: false, Message: vals}); err != nil {
		h.kit.ErrorJSON(w, err, 500)
	}
}
