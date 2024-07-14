package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	CRUD "bank/sql_go"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "fighting"
	dbname   = "bank"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach the database: %v\n", err)
	}

	fmt.Println("Successfully connected!")

	q := CRUD.New(db)
	ctx := context.Background()
	arg := CRUD.GetAccountParam{
		ID: 1,
	}
	/*
		arg := CRUD.CreateAccountParam{
			Owner:    "Alice",
			Balance:  1000,
			Currency: "USD",
		}*/
	account, err := q.GetAccount(ctx, arg)

	if err != nil {
		fmt.Printf("I am so bad!\n")
	}
	//fmt.Printf("Created account: %+v\n", account)
	fmt.Printf("Get account: %+v\n", account)
}
