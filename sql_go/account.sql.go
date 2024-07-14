package CRUD

import (
	"context"

	_ "github.com/lib/pq"
)

const Createaccountsql = `
    INSERT INTO account(
		Owner,
		Balance,
		Currency
	)VALUES(
	    $1,$2,$3
	)RETURNING id,owner,balance,currency,createdat
`

type CreateAccountParam struct {
	Owner    string
	Balance  int64
	Currency string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParam) (Account, error) {
	tmp := q.db.QueryRowContext(ctx, Createaccountsql, arg.Owner, arg.Balance, arg.Currency)
	var A Account
	err := tmp.Scan(
		&A.ID,
		&A.Owner,
		&A.Balance,
		&A.Currency,
		&A.Createdat,
	)

	return A, err
}

/*func getAccount(db *sql.DB, owner string) account {

}

func updateAccount(db *sql.DB, owner string, balance int) {

}

func deleteAccount(db *sql.DB, owner string) {

}
*/
