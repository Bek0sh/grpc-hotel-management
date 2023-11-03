package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("please check your secret password, you have provided: %s", secretKey)
	}

	return &JwtMaker{secretKey: secretKey}, nil

}

func (j *JwtMaker) CreateToken(id string, role string, dur time.Duration) (string, error) {
	payload := NewPayload(id, role, dur)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(j.secretKey))
}

func (j *JwtMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(
		token, &Payload{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrorExpired) {
			return nil, ErrorExpired
		}
		return nil, fmt.Errorf("failed to verify token, error: %v", err)
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, fmt.Errorf("failed to find payload, error: %v", err)
	}

	return payload, nil

}
