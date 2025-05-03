package helper

import (
	"crypto/ecdsa"
	"fmt"
	"gateway/internal/environment"
	"io"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// LoadPublicKey loads the public key from a file
func LoadPublicKey() (*ecdsa.PublicKey, error) {
	file, err := os.Open(environment.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open private key file %s: %w", environment.PublicKeyPath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("failed to close public key file: %w", err)
		}
	}(file)

	keyData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key file %s: %w", environment.PublicKeyPath, err)
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA public key from %s: %w", environment.PublicKeyPath, err)
	}

	return publicKey, nil
}
