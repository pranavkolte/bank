package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQuries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func TestGetUser(t *testing.T) {

	createdUser := createRandomUser(t)
	responseUser, err := testQuries.GetUser(context.Background(), createdUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, responseUser)

	require.Equal(t, createdUser.Username, responseUser.Username)
	require.Equal(t, createdUser.FullName, responseUser.FullName)
	require.Equal(t, createdUser.Email, responseUser.Email)
	require.Equal(t, createdUser.HashedPassword, responseUser.HashedPassword)

	require.WithinDuration(t, createdUser.PasswordChangedAt, responseUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, createdUser.CreatedAt, responseUser.CreatedAt, time.Second)

}
