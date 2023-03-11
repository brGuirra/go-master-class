package db

import (
	"context"
	"testing"
	"time"

	"github.com/brGuirra/simple-bank/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(gofakeit.Password(true, true, true, true, false, 8))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       gofakeit.Username(),
		HashedPassword: hashedPassword,
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
	gotUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, createdUser.Username, gotUser.Username)
	require.Equal(t, createdUser.HashedPassword, gotUser.HashedPassword)
	require.Equal(t, createdUser.FullName, gotUser.FullName)
	require.Equal(t, createdUser.Email, gotUser.Email)
	require.WithinDuration(t, createdUser.PasswordChangedAt, gotUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, createdUser.CreatedAt, gotUser.CreatedAt, time.Second)
}
