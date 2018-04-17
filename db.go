package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	fmt.Println("connecting to database")
	var err error
	Db, err = sql.Open("postgres", "user=seker dbname=seker_dev password=seker sslmode=disable")
	if err != nil {
		panic(err)
	}
}
