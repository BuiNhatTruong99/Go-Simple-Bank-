package db

import (
	"context"
	"github.com/BuiNhatTruong99/Go-Simple-Bank-/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
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
	entry := createRandomEntry(t, account)

	getEntry, err := testQueries.GetEntries(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntry)

	require.Equal(t, entry.ID, getEntry.ID)
	require.Equal(t, entry.AccountID, getEntry.AccountID)
	require.Equal(t, entry.Amount, getEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, getEntry.CreatedAt, time.Millisecond)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, arg.AccountID)
	}
}
