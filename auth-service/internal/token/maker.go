package token

import (
	"github.com/mohammadyaseen2/TokenUtils/model"
)

// TokenMaker is an interface for managing tokens
type TokenMaker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(*model.Payload) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(string) (*model.Payload, error)
}

// KeysMaker is an interface for managing key pair
type KeysMaker interface {
	// CreateKeyPair creates a new keys and save it in a specific files
	CreateKeyPair(string, string) (*model.KeyPair, error)

	// GetKeyPair get keys from a specific files
	GetKeyPair(string, string) (*model.KeyPair, error)
}
