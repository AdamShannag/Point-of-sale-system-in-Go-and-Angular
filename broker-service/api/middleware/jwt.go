package middleware

import (
	pb "broker-service/internal/grpc/proto/auth"
	"context"
	"errors"
	"log"
	"net/http"
)

func (m *Middleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			m.kit.ErrorJSON(w, errors.New("authentication failed, missing or invalid token"), http.StatusUnauthorized)
			return
		}

		if ok, err := m.isValidToken(w, token); err != nil {
			log.Println(err)
			m.kit.ErrorJSON(w, errors.New("authentication failed, missing or invalid token"), http.StatusUnauthorized)
		} else if ok {
			next.ServeHTTP(w, r)
		} else if !ok {
			m.kit.ErrorJSON(w, errors.New("invalid token"), http.StatusUnauthorized)
		}
	})
}

func (m *Middleware) isValidToken(w http.ResponseWriter, token string) (bool, error) {
	if res, err := m.authClient.VerifyToken(context.Background(), &pb.VerifyTokenRequest{Token: token}); err != nil {
		return false, err
	} else {
		return res.Verified, nil
	}
}
