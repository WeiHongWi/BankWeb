package main

import "time"

type account struct {
	ID        int64
	owner     string
	balance   int64
	currency  string
	createdat time.Time
}

type entries struct {
	ID        int64
	AccountID int64
	Amount    int64
	createdat time.Time
}

type transac struct {
	ID            int64
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	createdat     time.Time
}
