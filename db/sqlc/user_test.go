package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/duythien2212/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomOwner(),
		HashPassword: hashedPassword,
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

	require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreateAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashPassword, user2.HashPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.PasswordChangeAt, user2.PasswordChangeAt, time.Second)
	require.WithinDuration(t, user1.CreateAt, user2.CreateAt, time.Second)
}

func TestUpdateOnlyFullname(t *testing.T) {
	oldUser := createRandomUser(t)
	newFullName := util.RandomOwner()

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, newUser.FullName, oldUser.FullName)
	require.Equal(t, newUser.FullName, newFullName)
	require.Equal(t, newUser.HashPassword, oldUser.HashPassword)
	require.Equal(t, newUser.Email, oldUser.Email)
}

func TestUpdateOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)
	newEmail := util.RandomEmail()

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, newUser.Email, oldUser.Email)
	require.Equal(t, newUser.Email, newEmail)
	require.Equal(t, newUser.FullName, oldUser.FullName)
	require.Equal(t, newUser.HashPassword, oldUser.HashPassword)
}

func TestUpdateOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)
	newPassword := util.RandomString(6)
	newHashPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		HashPassword: sql.NullString{
			String: newHashPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, newUser.HashPassword, oldUser.HashPassword)
	require.Equal(t, newUser.HashPassword, newHashPassword)
	require.Equal(t, newUser.FullName, oldUser.FullName)
	require.Equal(t, newUser.Email, oldUser.Email)
}
