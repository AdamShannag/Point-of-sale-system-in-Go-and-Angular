package auth

import (
	"github.com/AdamShannag/toolkit"
	"github.com/go-chi/chi/v5"
)

type Auth struct {
	*chi.Mux
	kit *toolkit.Tools
}

func NewAuth(toolkit *toolkit.Tools) Auth {
	h := Auth{
		Mux: chi.NewMux(),
		kit: toolkit,
	}

	h.Post("/signup", h.Signup)
	h.Post("/signin", h.Signin)
	h.Post("/signout", h.Signout)
	h.Get("/signedin", h.Signedin)
	h.Post("/username", h.Username)

	return h
}
