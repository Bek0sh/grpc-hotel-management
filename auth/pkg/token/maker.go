package token

import "time"

type Maker interface {
	CreateToken(id string, role string, dur time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
