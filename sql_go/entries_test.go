package CRUD

import (
	"bank/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T) Entries {
	account := createRandomAcount(t)

	arg := CreateEntriesParam{
		Account_id: account.ID,
		Amount:     util.Random_money(),
	}

	row, err := test.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, row)

	require.Equal(t, arg.Account_id, row.Account_id)
	require.Equal(t, arg.Amount, row.Amount)

	require.NotZero(t, row.ID)
	require.NotZero(t, row.Createdat)

	return row
}

func TestCreateEntries(t *testing.T) {
	createRandomEntries(t)
}

func TestGetEntries(t *testing.T) {
	entry := createRandomEntries(t)

	arg := GetEntriesParam{
		ID: entry.ID,
	}

	tmp, err := test.GetEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tmp)

	require.Equal(t, tmp.ID, entry.ID)
	require.Equal(t, tmp.Account_id, entry.Account_id)
	require.Equal(t, tmp.Amount, entry.Amount)
	require.Equal(t, tmp.Createdat, entry.Createdat)

}

func TestListEntries(t *testing.T) {
	account := createRandomAcount(t)
	arg_count := CountEntriesParam{
		Account_id: account.ID,
	}
	offset, err := test.CountEntries(context.Background(), arg_count)
	require.NoError(t, err)

	arg := ListEntriesParam{
		Account_id: account.ID,
		Limit:      5,
		Offset:     int32(offset),
	}

	var entry_arr1 []Entries

	var i int
	for i = 0; i < int(arg.Limit); i++ {
		arg_e := CreateEntriesParam{
			Account_id: account.ID,
			Amount:     util.Random_money(),
		}

		tmp, err := test.CreateEntries(context.Background(), arg_e)

		require.NoError(t, err)
		require.NotEmpty(t, tmp)

		entry_arr1 = append(entry_arr1, tmp)
	}

	entry_arr2, err1 := test.ListEntries(context.Background(), arg)
	require.NoError(t, err1)

	var j int64
	for j = 0; j < int64(arg.Limit); j++ {
		require.NotEmpty(t, entry_arr2[j])
		require.Equal(t, entry_arr1[j].ID, entry_arr2[j].ID)
		require.Equal(t, entry_arr1[j].Amount, entry_arr2[j].Amount)
		require.Equal(t, entry_arr1[j].Createdat, entry_arr2[j].Createdat)
	}

}
