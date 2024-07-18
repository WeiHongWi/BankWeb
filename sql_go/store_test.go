package CRUD

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransactionTx(t *testing.T) {
	store := New_store(db)

	a1 := createRandomAcount(t)
	a2 := createRandomAcount(t)

	n := 5
	amount := int64(10)

	results := make(chan TransactionTxResult)
	errs := make(chan error)
	for i := 0; i < n; i++ {
		go func() {
			arg_transac := TransactionTxParam{
				From_account_id: a1.ID,
				To_account_id:   a2.ID,
				Amount:          amount,
			}

			result, err := store.TransactionTx(context.Background(), arg_transac)

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//Check the consistency of transaction.
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.From_account_id, a1.ID)
		require.Equal(t, transfer.To_account_id, a2.ID)
		require.Equal(t, transfer.Amount, amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.Createdat)

		arg_transac := GetTransacParam{
			ID: transfer.ID,
		}
		_, err = store.GetTransac(context.Background(), arg_transac)
		require.NoError(t, err)

		//Check the consistency of from account a1.
		from_entry := result.From_entry
		require.NotEmpty(t, from_entry)
		require.Equal(t, from_entry.Account_id, a1.ID)
		require.Equal(t, from_entry.Amount, -amount)
		require.NotZero(t, from_entry.ID)
		require.NotZero(t, from_entry.Createdat)

		arg_from := GetEntriesParam{
			ID: from_entry.ID,
		}
		_, err = store.GetEntries(context.Background(), arg_from)
		require.NoError(t, err)

		//Check the consistency of to account a2.
		to_entry := result.To_entry
		require.NotEmpty(t, to_entry)
		require.Equal(t, to_entry.Account_id, a2.ID)
		require.Equal(t, to_entry.Amount, amount)
		require.NotZero(t, to_entry.ID)
		require.NotZero(t, to_entry.Createdat)

		arg_to := GetEntriesParam{
			ID: to_entry.ID,
		}
		_, err = store.GetEntries(context.Background(), arg_to)
		require.NoError(t, err)

		//Check the from account
		from_account := result.From_account
		require.NotEmpty(t, from_account)
		require.Equal(t, a1.ID, from_account.ID)

		//Check the to account
		to_account := result.To_account
		require.NotEmpty(t, to_account)
		require.Equal(t, a2.ID, to_account.ID)

		//Check the account's balance
		dif_1 := a1.Balance - from_account.Balance
		dif_2 := to_account.Balance - a2.Balance
		require.Equal(t, dif_1, dif_2)
		require.True(t, dif_1 > 0)
		require.True(t, dif_1%amount == 0)
		k := int(dif_1 / amount)

		require.True(t, 1 <= k && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

	}
	//Check the final balance for both from account and to account
	update_a1, err := store.GetAccount(context.Background(), GetAccountParam{
		ID: a1.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, update_a1)
	require.Equal(t, update_a1.Balance+int64(n)*amount, a1.Balance)

	update_a2, err := store.GetAccount(context.Background(), GetAccountParam{
		ID: a2.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, update_a2)
	require.Equal(t, update_a2.Balance-int64(n)*amount, a2.Balance)

}
