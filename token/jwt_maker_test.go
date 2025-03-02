package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker("secretkeymoresecretthansize32sasaadasddasd")
	require.NoError(t, err)

	token, err := maker.CreateToken("some-username", time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestJWTMaker_VerifyToken(t *testing.T) {
	maker, err := NewJWTMaker("secretkeymoresecretthansize32sasaadasddasd")
	require.NoError(t, err)

	token, err := maker.CreateToken("some-username", time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	validUsername, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, "some-username", validUsername.Username)
}
