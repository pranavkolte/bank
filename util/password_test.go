package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	wrongpassword := RandomString(6)

	hashedpasssword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpasssword)

	err = CheckPassword(password, hashedpasssword)
	require.NoError(t, err)

	err = CheckPassword(wrongpassword, hashedpasssword)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword)

	hashedpassswordUnique, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassswordUnique)
	require.NotEqual(t, hashedpasssword, hashedpassswordUnique)
}
