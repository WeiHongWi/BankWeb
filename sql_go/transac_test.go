package CRUD

import (
	"bank/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransac(t *testing.T) Transac {
	account1 := createRandomAcount(t)
	account2 := createRandomAcount(t)

	arg := CreateTransacParam{
		Amount:          util.Random_money(),
		From_account_id: account1.ID,
		To_account_id:   account2.ID,
	}

	row, err := test.CreateTransac(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, row)

	require.Equal(t, row.From_account_id, arg.From_account_id)
	require.Equal(t, row.To_account_id, arg.To_account_id)
	require.Equal(t, row.Amount, arg.Amount)

	require.NotZero(t, row.ID)
	require.NotZero(t, row.Createdat)

	return row
}
func TestCreateTransac(t *testing.T) {
	createRandomTransac(t)
}

func TestGetTransac(t *testing.T) {
	transac := createRandomTransac(t)

	arg := GetTransacParam{
		ID: transac.ID,
	}

	row, err := test.GetTransac(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, row)

	require.Equal(t, row.ID, transac.ID)
	require.Equal(t, row.From_account_id, transac.From_account_id)
	require.Equal(t, row.To_account_id, transac.To_account_id)
	require.Equal(t, row.Amount, transac.Amount)
}

func TestListTransac(t *testing.T) {
	a1 := createRandomAcount(t)
	a2 := createRandomAcount(t)

	arg_count := CountTransacParam{
		From_account_id: a1.ID,
		To_account_id:   a2.ID,
	}
	count, err := test.CountTransac(context.Background(), arg_count)

	require.Error(t, err)

	arg := ListTransacParam{
		From_account_id: a1.ID,
		To_account_id:   a2.ID,
		Limit:           5,
		Offset:          int32(count),
	}

	var transac_arr []Transac

	var i int64
	for i = 0; i < int64(arg.Limit); i++ {
		arg_create := CreateTransacParam{
			From_account_id: a1.ID,
			To_account_id:   a2.ID,
			Amount:          util.Random_money(),
		}

		row, err := test.CreateTransac(context.Background(), arg_create)

		require.NoError(t, err)
		require.NotEmpty(t, row)

		transac_arr = append(transac_arr, row)
	}

	transac_arr1, err := test.ListTransac(context.Background(), arg)

	require.NoError(t, err)

	var j int64
	for j = 0; j < int64(arg.Limit); j++ {
		require.NotEmpty(t, transac_arr1[j])
		require.Equal(t, transac_arr1[j].ID, transac_arr[j].ID)
		require.Equal(t, transac_arr1[j].Amount, transac_arr[j].Amount)
		require.Equal(t, transac_arr1[j].From_account_id, transac_arr[j].From_account_id)
		require.Equal(t, transac_arr1[j].To_account_id, transac_arr[j].To_account_id)
	}
}
