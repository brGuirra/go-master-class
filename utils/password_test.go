package utils

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func generateRandomPassword(t *testing.T) (password string, hashedPassword string) {
	var err error
	password = gofakeit.Password(true, true, true, true, false, 8)
	hashedPassword, err = HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	return
}

func TestHashPassword(t *testing.T) {
	password, hashedPassword1 := generateRandomPassword(t)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

func TestCheckPassword(t *testing.T) {
	t.Run("valid password", func(t *testing.T) {
		password, hashedPassword := generateRandomPassword(t)
		err := CheckPassword(password, hashedPassword)
		require.NoError(t, err)
	})

	t.Run("wrong password", func(t *testing.T) {
		_, hashedPassword := generateRandomPassword(t)
		wrongPassword := gofakeit.Password(true, true, true, true, false, 8)

		err := CheckPassword(wrongPassword, hashedPassword)
		require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	})
}
