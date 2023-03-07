package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       gofakeit.Username(),
		HashedPassword: gofakeit.Password(true, true, true, true, false, 8),
		FullName:       gofakeit.Name(),
		Email:          gofakeit.Email(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.Zero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	gettedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, gettedUser)

	require.Equal(t, createdUser.Username, gettedUser.Username)
	require.Equal(t, createdUser.HashedPassword, gettedUser.HashedPassword)
	require.Equal(t, createdUser.FullName, gettedUser.FullName)
	require.Equal(t, createdUser.Email, gettedUser.Email)
	require.WithinDuration(t, createdUser.PasswordChangedAt, gettedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, createdUser.CreatedAt, gettedUser.CreatedAt, time.Second)
}
