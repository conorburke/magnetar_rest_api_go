package main

import (
	"fmt"
)

func allUsers() (users []user, err error) {
	rows, err := Db.Query("select first_name, last_name, email from users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := user{}
		err = rows.Scan(&user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return
		}
		fmt.Println(user)
		users = append(users, user)
		fmt.Println(users)
	}
	rows.Close()
	return
}
