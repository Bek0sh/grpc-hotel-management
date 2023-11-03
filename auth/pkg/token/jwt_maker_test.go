package token_test

import (
	"testing"
	"time"

	"github.com/Bek0sh/hotel-management-auth/pkg/token"
	"github.com/stretchr/testify/require"
)

func TestJwtToken(t *testing.T) {
	maker, err := token.NewJwtMaker("fdsafsdjifejrifncjsdnajhfuejrinvajierunjgjiafiernvjnsfdugeiar")
	require.NoError(t, err)

	// issAt := time.Now()
	duration := 1 * time.Minute
	// exp := issAt.Add(duration)

	tok, err := maker.CreateToken("2", "user", duration)

	require.NoError(t, err)
	require.NotEmpty(t, tok)

	payload, err := maker.VerifyToken(tok)
	require.NoError(t, err)

	require.Equal(t, payload.Id, "2")

}

func TestExpiredToken(t *testing.T) {
	maker, err := token.NewJwtMaker("fdsafsdjifejrifncjsdnajhfuejrinvajierunjgjiafiernvjnsfdugeiar")
	require.NoError(t, err)

	dur := 10 * time.Second

	tok, err := maker.CreateToken("1", "user", dur)
	require.NoError(t, err)

	time.Sleep(15 * time.Second)

	payload, err := maker.VerifyToken(tok)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrorExpired.Error())
	require.Nil(t, payload)
}
