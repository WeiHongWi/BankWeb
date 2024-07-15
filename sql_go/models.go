package CRUD

import "time"

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	Createdat time.Time
}

type Entries struct {
	ID         int64
	Account_id int64
	Amount     int64
	Createdat  time.Time
}

type Transac struct {
	ID            int64
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	Createdat     time.Time
}
