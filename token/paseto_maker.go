package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a struct that implements the Maker interface using the PASETO protocol.
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker is an interface for creating and verifying tokens.
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid symmetric key length: must be %d bytes, got %d bytes", chacha20poly1305.KeySize, len(symmetricKey))
	}

	paseto := paseto.NewV2()
	maker := &PasetoMaker{
		paseto:       paseto,
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// CreateToken create paseto token
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken verify paseto token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if err := payload.Valid(); err != nil {
		return nil, fmt.Errorf("invalid payload: %w", err)
	}

	return payload, nil
}
