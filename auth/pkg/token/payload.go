package token

import (
	"fmt"
	"time"
)

var (
	ErrorExpired = fmt.Errorf("You token have expired")
)

type Payload struct {
	Id        string
	UserRole  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

func NewPayload(id string, role string, dur time.Duration) *Payload {
	now := time.Now()
	exp := time.Now().Add(dur)

	payload := &Payload{
		Id:        id,
		UserRole:  role,
		IssuedAt:  now,
		ExpiresAt: exp,
	}

	return payload
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return ErrorExpired
	}
	return nil
}
