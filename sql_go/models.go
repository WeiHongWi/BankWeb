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
	ID              int64
	From_account_id int64
	To_account_id   int64
	Amount          int64
	Createdat       time.Time
}
