package token

import (
	b64 "encoding/base64"
	"github.com/mohammadyaseen2/TokenUtils/jwt"
	"github.com/mohammadyaseen2/TokenUtils/model"
	"log"
	"os"
)

type KeyPairMaker struct {
	bits int
}

func NewKeysMaker(bits int) KeysMaker {
	return &KeyPairMaker{bits}
}

func (maker *KeyPairMaker) CreateKeyPair(privateKeysPath, publicKeysPath string) (*model.KeyPair, error) {
	keys, err := jwt.MakeKeyPair(maker.bits)

	privateKey, publicKey, done := createKeysFile(privateKeysPath, publicKeysPath)

	defer privateKey.Close()
	defer publicKey.Close()

	if done {
		saveKeysFile(keys, err, privateKey, publicKey)
	}

	return keys, err
}

func (maker *KeyPairMaker) GetKeyPair(privateKeysPath, publicKeysPath string) (*model.KeyPair, error) {
	privateKey, publicKey, err := getKeys(privateKeysPath, publicKeysPath)
	if err != nil {
		return nil, err
	}

	return &model.KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func decodeKeys(privateKey []byte, publicKey []byte) ([]byte, []byte, error) {
	decodedPrivateKey, err := b64.StdEncoding.DecodeString(string(privateKey))
	if err != nil {
		return nil, nil, err
	}

	decodedPublicKey, err := b64.StdEncoding.DecodeString(string(publicKey))
	if err != nil {
		return nil, nil, err
	}
	return decodedPrivateKey, decodedPublicKey, nil
}

func getKeys(privateKeysPath string, publicKeysPath string) ([]byte, []byte, error) {
	privateKey, err := os.ReadFile(privateKeysPath)
	if err != nil {
		return nil, nil, err
	}
	publicKey, err := os.ReadFile(publicKeysPath)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, publicKey, nil
}

func saveKeysFile(keys *model.KeyPair, err error, privateKey *os.File, publicKey *os.File) {
	encodedPrivateKey := b64.StdEncoding.EncodeToString(keys.PrivateKey)
	encodedPublicKey := b64.StdEncoding.EncodeToString(keys.PublicKey)
	_, err = privateKey.WriteString(encodedPrivateKey)
	_, err = publicKey.WriteString(encodedPublicKey)

	log.Println("save key successfully")
}

func createKeysFile(privateKeysPath string, publicKeysPath string) (*os.File, *os.File, bool) {
	done := false

	privateKey, err := os.Create(privateKeysPath)

	if err != nil {
		done = true
		log.Printf("error while save private key in file path [%s]: %s\n", privateKeysPath, err)
	}

	publicKey, err := os.Create(publicKeysPath)

	if err != nil {
		done = true
		log.Printf("error while save public key in file path [%s]: %s\n", publicKeysPath, err)

	}
	return privateKey, publicKey, done
}
