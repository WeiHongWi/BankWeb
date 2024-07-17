package CRUD

import (
	"bank/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAcount(t *testing.T) Account {
	/*
		var arg = CreateAccountParam{
			Owner:    "Willy",
			Balance:  100,
			Currency: "NT",
		}*/

	var arg = CreateAccountParam{
		Owner:    util.Random_owner(),
		Balance:  util.Random_money(),
		Currency: util.Random_currency(),
	}

	account, err := test.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.Createdat)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAcount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAcount(t)
	arg := GetAccountParam{
		ID: account1.ID,
	}
	account2, err := test.GetAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAcount(t)
	arg := UpdateAccountParam{
		Balance: 500,
		ID:      account1.ID,
	}

	arg_get := GetAccountParam{
		ID: arg.ID,
	}

	err := test.UpdateAccount(context.Background(), arg)
	account2, err2 := test.GetAccount(context.Background(), arg_get)

	require.NoError(t, err)
	require.NoError(t, err2)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
}

func TestListAccount(t *testing.T) {

	var len_account int64 = test.CountOfAccount(context.Background())
	arg := ListAccountParam{
		Limit:  5,
		Offset: int32(len_account),
	}

	var i int
	var account_list1 []Account

	for i = 0; i < int(arg.Limit); i++ {
		a := createRandomAcount(t)
		account_list1 = append(account_list1, a)
	}

	account_list2, err := test.ListAccount(context.Background(), arg)
	require.NoError(t, err)

	var j int
	for j = 0; j < int(arg.Limit); j++ {
		require.NotEmpty(t, account_list2[j])
		require.Equal(t, account_list1[j].ID, account_list2[j].ID)
		require.Equal(t, account_list1[j].Balance, account_list2[j].Balance)
		require.Equal(t, account_list1[j].Currency, account_list2[j].Currency)
		require.Equal(t, account_list1[j].Owner, account_list2[j].Owner)
	}

}
