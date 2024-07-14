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
	var arg CRUD.CreateAccountParam

	arg.Owner = `HongWei`
	arg.Balance = 100
	arg.Currency = `NT`

	var row CRUD.Account
	row, err = (*CRUD.Queries).CreateAccount(q, context.Background(), arg)

	if err != nil {
		fmt.Printf("I am so bad!\n")
	}
	fmt.Printf("ID is %d\n", row.ID)

}
