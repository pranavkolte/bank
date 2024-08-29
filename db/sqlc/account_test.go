package db

import (
	"bank/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQuries.CreateAccount(context.Background(), arg)

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
	// create account - should be independent

	createdAccount := createRandomAccount(t)
	responseAccount, err := testQuries.GetAccount(context.Background(), createdAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, responseAccount)

	require.Equal(t, createdAccount.ID, responseAccount.ID)
	require.Equal(t, createdAccount.Owner, responseAccount.Owner)
	require.Equal(t, createdAccount.Balance, responseAccount.Balance)
	require.Equal(t, createdAccount.Currency, responseAccount.Currency)

	require.WithinDuration(t, createdAccount.CreatedAt, responseAccount.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQuries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, createdAccount.ID, updatedAccount.ID)
	require.Equal(t, createdAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, createdAccount.Currency, updatedAccount.Currency)

	require.WithinDuration(t, createdAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	err := testQuries.DeleteAccount(context.Background(), createdAccount.ID)

	require.NoError(t, err)

	deletedAccount, err := testQuries.GetAccount(context.Background(), createdAccount.ID)

	require.Empty(t, deletedAccount)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	// var lastAccount Account
	for i := 0; i < 10; i++ {
		// lastAccount = createRandomAccount(t)
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQuries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
