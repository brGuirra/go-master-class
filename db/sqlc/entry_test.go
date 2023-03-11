package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    int64(gofakeit.Number(100, 1000)),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	createdEntry := createRandomEntry(t, account)
	gotEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotEntry)

	require.Equal(t, createdEntry.ID, gotEntry.ID)
	require.Equal(t, createdEntry.AccountID, gotEntry.AccountID)
	require.Equal(t, createdEntry.Amount, gotEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, gotEntry.CreatedAt, time.Second)
}
