package token

import (
	"github.com/mohammadyaseen2/TokenUtils/jwt"
	"github.com/mohammadyaseen2/TokenUtils/model"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	jwt       jwt.JWT
	validator jwt.JWTValidator
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(jwt jwt.JWT, validator jwt.JWTValidator) *JWTMaker {
	return &JWTMaker{jwt, validator}
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(payload *model.Payload) (string, error) {
	return maker.jwt.GenerateToken(payload)
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*model.Payload, error) {
	return maker.validator.ValidateToken(token)
}
