package main

import (
	"bank/api"
	CRUD "bank/sql_go"
	"database/sql"
	"fmt"
	"log"
)

const (
	host          = "localhost"
	port          = 5432
	user          = "root"
	password      = "fighting"
	dbname        = "bank"
	serveraddress = "0.0.0.0:8080"
)

var db *sql.DB

func main() {
	DBinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", DBinfo)
	store := CRUD.New_store(db)
	server := api.NewServer(store)

	err = server.Start(serveraddress)
	if err != nil {
		log.Printf("Cannot start the server! %s\n", err)
	}
}
