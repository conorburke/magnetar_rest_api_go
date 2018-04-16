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

func (user *user) createUser() (err error) {
	statement := "insert into users (first_name, last_name, email) values ($1, $2, $3) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email).Scan(&user.Id)
	return
}
