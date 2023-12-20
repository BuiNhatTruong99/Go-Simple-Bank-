package db

import (
	"context"
	"github.com/BuiNhatTruong99/Go-Simple-Bank-/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	hashPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomOwner(),
		HashPassword: hashPassword,
		FullName:     util.RandomOwner(),
		Email:        util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashPassword, user.HashPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangeAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	getuser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, getuser)

	require.Equal(t, getuser.Username, user.Username)
	require.Equal(t, getuser.HashPassword, user.HashPassword)
	require.Equal(t, getuser.FullName, user.FullName)
	require.Equal(t, getuser.Email, user.Email)
	require.WithinDuration(t, user.CreatedAt, getuser.CreatedAt, time.Millisecond)
	require.WithinDuration(t, user.PasswordChangeAt, getuser.PasswordChangeAt, time.Millisecond)

}
