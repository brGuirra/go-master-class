package db

import (
	"context"
	"database/sql"
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
	gettedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gettedEntry)

	require.Equal(t, createdEntry.ID, gettedEntry.ID)
	require.Equal(t, createdEntry.AccountID, gettedEntry.AccountID)
	require.Equal(t, createdEntry.Amount, gettedEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, gettedEntry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createdEntry := createRandomEntry(t, account)

	arg := UpdateEntryParams{
		ID:     createdEntry.ID,
		Amount: int64(gofakeit.Number(100, 1000)),
	}
	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, createdEntry.ID, updatedEntry.ID)
	require.Equal(t, createdEntry.AccountID, updatedEntry.AccountID)
	require.Equal(t, arg.Amount, updatedEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, updatedEntry.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	createdEntry := createRandomEntry(t, account)

	err := testQueries.DeleteEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)

	gettedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, gettedEntry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		account := createRandomAccount(t)
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
