package CRUD

import (
	"context"

	_ "github.com/lib/pq"
)

const CreateEntriesSQL = `
INSERT INTO "Entries"
("Account_id","Amount")
VALUES
($1,$2)
RETURNING "ID","Account_id","Amount","Createdat"
`

type CreateEntriesParam struct {
	Account_id int64
	Amount     int64
}

const GetEntriesSQL = `
SELECT 
"ID","Account_id","Amount","Createdat"
FROM "Entries" WHERE "ID" = $1
LIMIT 1`

type GetEntriesParam struct {
	ID int64
}

const ListEntriesSQL = `
SELECT 
"ID","Account_id","Amount","Createdat"
FROM "Entries"
WHERE "Account_id" = $1
ORDER BY "ID"
LIMIT $2
OFFSET $3
`

type ListEntriesParam struct {
	Account_id int64
	Limit      int32
	Offset     int32
}

const CountEntriesSQL = `
SELECT
COUNT(*)
FROM "Entries"
WHERE "Account_id" = $1
`

type CountEntriesParam struct {
	Account_id int64
}

func (q *Queries) CountEntries(ctx context.Context, arg CountEntriesParam) (int64, error) {
	tmp := q.db.QueryRowContext(ctx, CountEntriesSQL, arg.Account_id)

	var count int64
	err := tmp.Scan(&count)

	return count, err
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParam) (Entries, error) {
	tmp := q.db.QueryRowContext(ctx, CreateEntriesSQL, arg.Account_id, arg.Amount)

	var E Entries
	err := tmp.Scan(
		&E.ID,
		&E.Account_id,
		&E.Amount,
		&E.Createdat,
	)

	return E, err
}

func (q *Queries) GetEntries(ctx context.Context, arg GetEntriesParam) (Entries, error) {
	tmp := q.db.QueryRowContext(ctx, GetEntriesSQL, arg.ID)

	var E Entries
	err := tmp.Scan(
		&E.ID,
		&E.Account_id,
		&E.Amount,
		&E.Createdat,
	)

	return E, err
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParam) ([]Entries, error) {
	rows, err := q.db.QueryContext(ctx, ListEntriesSQL, arg.Account_id, arg.Limit, arg.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var E_arr []Entries

	for rows.Next() {
		var E Entries
		if err = rows.Scan(
			&E.ID,
			&E.Account_id,
			&E.Amount,
			&E.Createdat,
		); err != nil {
			return nil, err
		}
		E_arr = append(E_arr, E)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return E_arr, nil
}
