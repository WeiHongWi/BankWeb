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

			results <- result
			errs <- err
		}()
	}

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
	}

}
