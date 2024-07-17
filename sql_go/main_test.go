package CRUD

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "fighting"
	dbname   = "bank"
)

var test *Queries
var db *sql.DB

func TestMain(m *testing.M) {
	DBinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", DBinfo)

	if err != nil {
		log.Fatalln("Open database failed!")
	}

	test = New(db)

	os.Exit(m.Run())

}
