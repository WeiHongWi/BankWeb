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

func (store *Store) execStore(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	q := New(tx)

	if err != nil {
		return err
	}

	err = fn(q)
	if err != nil {
		roll_err := tx.Rollback()
		if roll_err != nil {
			err = fmt.Errorf("roll err is:%v,err is: %v", roll_err, err)
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

type TransactionParam struct {
	From_account_id int64
	To_account_id   int64
	Amount          int64
}

type TransactionResult struct {
	transac            Transac
	from_account_entry Entries
	to_account_entry   Entries
	from_account       Account
	to_account         Account
}

func (store *Store) Transaction(ctx context.Context, arg TransactionParam) (TransactionResult, error) {
	var result TransactionResult

	err := store.execStore(ctx, func(q *Queries) error {
		var transac_err error
		result.transac, transac_err = q.CreateTransac(ctx, CreateTransacParam{
			From_account_id: arg.From_account_id,
			To_account_id:   arg.To_account_id,
			Amount:          arg.Amount,
		})

		if transac_err != nil {
			return transac_err
		}

		var from_entry_err error
		result.from_account_entry, from_entry_err = q.CreateEntries(ctx, CreateEntriesParam{
			Account_id: arg.From_account_id,
			Amount:     -arg.Amount,
		})
		if from_entry_err != nil {
			return from_entry_err
		}

		var to_entry_err error
		result.to_account_entry, to_entry_err = q.CreateEntries(ctx, CreateEntriesParam{
			Account_id: arg.To_account_id,
			Amount:     arg.Amount,
		})
		if to_entry_err != nil {
			return to_entry_err
		}
		return nil
	})

	return result, err
}
