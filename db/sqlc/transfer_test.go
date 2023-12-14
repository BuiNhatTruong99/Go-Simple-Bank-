package db

import (
	"GoSimpleBank/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)
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
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, account1, account2)

	getTransfer, err := testQueries.GetTransfers(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)

	require.Equal(t, transfer.ID, getTransfer.ID)
	require.Equal(t, transfer.FromAccountID, getTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, getTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, getTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, getTransfer.CreatedAt, time.Millisecond)
}

func TestGetListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
