package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type entries struct {
	account_id int
	amount     int
}

func main() {
	// 连接数据库
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
}

func createEntries(db *sql.DB, account_id int, amount int) {

}

func getEntries(db *sql.DB, owner string) entries {

}

func updateEntries(db *sql.DB, account_id int, amount int) {

}

func deleteEntries(db *sql.DB, account_id int) {

}
