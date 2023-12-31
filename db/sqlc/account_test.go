package db

import (
	"context"
	"database/sql"
	"github.com/BuiNhatTruong99/Go-Simple-Bank-/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	getAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, account.ID, getAccount.ID)
	require.Equal(t, account.Owner, getAccount.Owner)
	require.Equal(t, account.Balance, getAccount.Balance)
	require.Equal(t, account.Currency, getAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, getAccount.CreatedAt, time.Millisecond)
}

func TestGetListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	currentAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      currentAccount.ID,
		Balance: util.RandomMoney(),
	}

	updateAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount)

	require.Equal(t, currentAccount.ID, updateAccount.ID)
	require.Equal(t, currentAccount.Owner, updateAccount.Owner)
	require.Equal(t, arg.Balance, updateAccount.Balance)
	require.Equal(t, currentAccount.Currency, updateAccount.Currency)
	require.WithinDuration(t, currentAccount.CreatedAt, updateAccount.CreatedAt, time.Millisecond)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	getDeletedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getDeletedAccount)
}
