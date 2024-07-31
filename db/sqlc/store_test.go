package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	// running concurrent transactions
	n := 5
	amount := int64(50)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {

			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: from_account.ID,
				ToAccountID:   to_account.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// verify transfers
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, from_account.ID, transfer.FromAccountID)
		require.Equal(t, to_account.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// get from entry data
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)
		//check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, from_account.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		//get to entry
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, to_account.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// TODO: Check account balance
		// check accounts
		fromAccount := result.FomAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, from_account.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, to_account.ID, toAccount.ID)

		// check balances
		diff1 := from_account.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - to_account.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balance
	updatedAccount1, err := testQuries.GetAccountForUpdate(context.Background(), from_account.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQuries.GetAccountForUpdate(context.Background(), to_account.ID)
	require.NoError(t, err)

	require.Equal(t, from_account.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, to_account.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	// running concurrent transactions
	n := 10
	amount := int64(50)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountIDdead := from_account.ID
		toAccountIDdead := to_account.ID

		if i%2 == 1 {
			fromAccountIDdead = to_account.ID
			toAccountIDdead = from_account.ID
		}
		go func() {

			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountIDdead,
				ToAccountID:   toAccountIDdead,
				Amount:        amount,
			})

			errs <- err

		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}

	// check the final updated balance
	updatedAccount1, err := testQuries.GetAccountForUpdate(context.Background(), from_account.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQuries.GetAccountForUpdate(context.Background(), to_account.ID)
	require.NoError(t, err)

	require.Equal(t, from_account.Balance, updatedAccount1.Balance)
	require.Equal(t, to_account.Balance, updatedAccount2.Balance)
}
