package CRUD

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

const CreateAccountSQL = `INSERT INTO "Account" 
("Owner","Balance","Currency") 
VALUES($1,$2,$3) 
RETURNING "ID","Owner","Balance","Currency","Createdat"`

type CreateAccountParam struct {
	Owner    string
	Balance  int64
	Currency string
}

const GetAccountSQL = `SELECT 
					  	"ID","Owner","Balance","Currency","Createdat" 
					   FROM "Account" WHERE "ID" = $1`

type GetAccountParam struct {
	ID int64
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParam) (Account, error) {
	tmp := q.db.QueryRowContext(ctx, CreateAccountSQL, arg.Owner, arg.Balance, arg.Currency)
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

func (q *Queries) GetAccount(ctx context.Context, arg GetAccountParam) (Account, error) {
	tmp := q.db.QueryRowContext(ctx, GetAccountSQL, arg.ID)
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

func UpdateAccount(db *sql.DB, owner string, balance int) {

}

func DeleteAccount(db *sql.DB, owner string) {

}
