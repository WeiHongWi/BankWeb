package CRUD

import (
	"context"
)

const CountTransacSQL = `
SELECT
COUNT(*)
FROM "Transac"
WHERE "From_account_id" = $1 and
"To_account_id" = $2
`

type CountTransacParam struct {
	From_account_id int64
	To_account_id   int64
}

func (q *Queries) CountTransac(ctx context.Context, arg CountTransacParam) (int64, error) {
	tmp := q.db.QueryRowContext(ctx, CountTransacSQL)

	var count int64
	err := tmp.Scan(&count)

	return count, err
}

const CreateTransacSQL = `
INSERT INTO "Transac"
("From_account_id","To_account_id","Amount")
VALUES
($1,$2,$3)
RETURNING "ID","From_account_id","To_account_id","Amount","Createdat"
`

type CreateTransacParam struct {
	From_account_id int64
	To_account_id   int64
	Amount          int64
}

func (q *Queries) CreateTransac(ctx context.Context, arg CreateTransacParam) (Transac, error) {
	row := q.db.QueryRowContext(ctx, CreateTransacSQL, arg.From_account_id, arg.To_account_id, arg.Amount)

	var T Transac
	err := row.Scan(
		&T.ID,
		&T.From_account_id,
		&T.To_account_id,
		&T.Amount,
		&T.Createdat,
	)
	return T, err
}

const GetTransacSQL = `
SELECT "ID","From_account_id","To_account_id","Amount","Createdat"
FROM "Transac"
WHERE "ID" = $1
LIMIT 1
`

type GetTransacParam struct {
	ID int64
}

func (q *Queries) GetTransac(ctx context.Context, arg GetTransacParam) (Transac, error) {
	row := q.db.QueryRowContext(ctx, GetTransacSQL, arg.ID)

	var T Transac
	err := row.Scan(
		&T.ID,
		&T.From_account_id,
		&T.To_account_id,
		&T.Amount,
		&T.Createdat,
	)
	return T, err
}

const ListTransacSQL = `
SELECT "ID","From_account_id","To_account_id","Amount","Createdat"
FROM "Transac"
WHERE "From_account_id" = $1 OR
      "To_account_id" = $2
ORDER BY "ID"
LIMIT $3
OFFSET $4
`

type ListTransacParam struct {
	From_account_id int64
	To_account_id   int64
	Limit           int32
	Offset          int32
}

func (q *Queries) ListTransac(ctx context.Context, arg ListTransacParam) ([]Transac, error) {
	rows, err := q.db.QueryContext(ctx, ListTransacSQL, arg.From_account_id, arg.To_account_id, arg.Limit, arg.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var T_arr []Transac

	for rows.Next() {
		var T Transac
		if err := rows.Scan(
			&T.ID,
			&T.From_account_id,
			&T.To_account_id,
			&T.Amount,
			&T.Createdat,
		); err != nil {
			return nil, err
		}
		T_arr = append(T_arr, T)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return T_arr, nil
}
