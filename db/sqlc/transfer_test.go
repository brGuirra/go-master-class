package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from, to Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Amount:        int64(gofakeit.Number(100, 1000)),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	createRandomTransfer(t, from, to)
}

func TestGetTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	createdTransfer := createRandomTransfer(t, from, to)
	gotTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotTransfer)

	require.Equal(t, createdTransfer.ID, gotTransfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, gotTransfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, gotTransfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, gotTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, gotTransfer.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	createdTransfer := createRandomTransfer(t, from, to)

	arg := UpdateTransferParams{
		ID:     createdTransfer.ID,
		Amount: int64(gofakeit.Number(100, 1000)),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, createdTransfer.ID, updatedTransfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, updatedTransfer.ToAccountID)
	require.Equal(t, arg.Amount, updatedTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, updatedTransfer.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	createdTransfer := createRandomTransfer(t, from, to)

	err := testQueries.DeleteTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)

	gotTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, gotTransfer)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		from := createRandomAccount(t)
		to := createRandomAccount(t)
		createRandomTransfer(t, from, to)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
