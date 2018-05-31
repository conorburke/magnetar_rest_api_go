package main

import (
	"fmt"
)

func selectUsers() (users []user, err error) {
	rows, err := Db.Query("select * from users")
	if err != nil {
		return
	}
	for rows.Next() {
		u := user{}
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash,
			&u.PhoneNumber, &u.Age, &u.Address1, &u.Address2, &u.City, &u.Region, &u.Zipcode)
		if err != nil {
			return
		}
		fmt.Println(u)
		users = append(users, u)
		fmt.Println(users)
	}
	rows.Close()
	return
}

func selectUser(id int) (u user, err error) {
	u = user{}
	err = Db.QueryRow(`select * from users where id = $1`, id).
		Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash,
		&u.PhoneNumber, &u.Age, &u.Address1, &u.Address2, &u.City, &u.Region, &u.Zipcode)
	return
}

func insertUser(u *user) (id int, err error) {
	statement := "insert into users (first_name, last_name, email, phone_number) values ($1, $2, $3, $4) returning id"
	stmt, err := Db.Prepare(statement)
	fmt.Println(stmt)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.FirstName, u.LastName, u.Email, u.PhoneNumber).Scan(&u.ID)
	if err != nil {
		return
	}
	fmt.Println(u.ID)
	return u.ID, err
}

func destroyTool(id int) {
	Db.QueryRow(`delete from tools where id = $1`, id)
}

func (user *user) createUser() (err error) {
	statement := "insert into users (first_name, last_name, email) values ($1, $2, $3) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email).Scan(&user.ID)
	return
}

func selectTools() (tools []tool, err error) {
	rows, err := Db.Query("select * from tools")
	if err != nil {
		return
	}
	for rows.Next() {
		t := tool{}
		err = rows.Scan(&t.ID, &t.Title, &t.ToolType, &t.Price, &t.ToolOwner)
		if err != nil {
			return
		}
		fmt.Println(t)
		tools = append(tools, t)
		fmt.Println(tools)
	}
	rows.Close()
	return
}

func selectTool(id int) (t tool, err error) {
	t = tool{}
	err = Db.QueryRow(`select * from tools where id = $1`, id).
		Scan(&t.ID, &t.Title, &t.ToolType, &t.Price, &t.ToolOwner)
	return
}

func selectUserTools(id int) (tools []tool, err error) {
	rows, err := Db.Query("select * from tools where tool_owner = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		t := tool{}
		err = rows.Scan(&t.ID, &t.Title, &t.ToolType, &t.Price, &t.ToolOwner)
		if err != nil {
			return
		}
		fmt.Println(t)
		tools = append(tools, t)
		fmt.Println(tools)
	}
	rows.Close()
	return
}

func selectUserTool(uID int, tID int) (t tool, err error) {
	t = tool{}
	err = Db.QueryRow(`select * from tools where tool_owner = $1 and id = $2`, uID, tID).
		Scan(&t.ID, &t.Title, &t.ToolType, &t.Price, &t.ToolOwner)
	return
}