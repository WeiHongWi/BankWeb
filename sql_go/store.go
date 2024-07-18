package CRUD

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func New_store(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		roll_err := tx.Rollback()
		if roll_err != nil {
			return fmt.Errorf("roll err is:%v,err is: %v", roll_err, err)
		}
		return err
	}

	return tx.Commit()
}

//Implement the database transaction for transaction between two accounts
//There is 5 functions should invoke:
// (1) CreateTransac
// (2) CreateEntry for From_account_id
// (3) CreateEntry for To_account_id
// (4) Update Balance for From_account_id
// (5) Update Balance for To_account_id

type TransactionTxParam struct {
	From_account_id int64
	To_account_id   int64
	Amount          int64
}

type TransactionTxResult struct {
	Transfer     Transac
	From_entry   Entries
	To_entry     Entries
	From_account Account
	To_account   Account
}

func (store *Store) TransactionTx(ctx context.Context, arg TransactionTxParam) (TransactionTxResult, error) {
	var result TransactionTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransac(ctx, CreateTransacParam{
			From_account_id: arg.From_account_id,
			To_account_id:   arg.To_account_id,
			Amount:          arg.Amount,
		})

		if err != nil {
			return err
		}

		result.From_entry, err = q.CreateEntries(ctx, CreateEntriesParam{
			Account_id: arg.From_account_id,
			Amount:     -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.To_entry, err = q.CreateEntries(ctx, CreateEntriesParam{
			Account_id: arg.To_account_id,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}

		account1, err := q.GetAccountForUpdate(ctx, GetAccountForUpdateParam{
			ID: arg.From_account_id,
		})
		if err != nil {
			return err
		}

		err = q.UpdateAccount(ctx, UpdateAccountParam{
			Balance: account1.Balance - arg.Amount,
			ID:      account1.ID,
		})
		if err != nil {
			return err
		}

		result.From_account, err = q.GetAccountForUpdate(ctx, GetAccountForUpdateParam{
			ID: account1.ID,
		})
		if err != nil {
			return err
		}

		account2, err := q.GetAccount(ctx, GetAccountParam{
			ID: arg.To_account_id,
		})
		if err != nil {
			return err
		}

		err = q.UpdateAccount(ctx, UpdateAccountParam{
			Balance: account2.Balance + arg.Amount,
			ID:      account2.ID,
		})
		if err != nil {
			return err
		}

		result.To_account, err = q.GetAccount(ctx, GetAccountParam{
			ID: account2.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
