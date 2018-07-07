package main

import (
	"fmt"
)

func selectUsers() (users []user, err error) {
	rows, err := Db.Query(`select id, first_name, last_name, email, password_hash,
		phone_number, address_1, address_2, city, region, zipcode from users`)
	if err != nil {
		return
	}
	for rows.Next() {
		u := user{}
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash,
			&u.PhoneNumber, &u.Address1, &u.Address2, &u.City, &u.Region, &u.Zipcode)
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
	err = Db.QueryRow(`select id, first_name, last_name, email, password_hash,
		  phone_number, address_1, address_2, city, region, zipcode from users where id = $1`, id).
		Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash,
		&u.PhoneNumber, &u.Address1, &u.Address2, &u.City, &u.Region, &u.Zipcode)
	return
}

func insertTool(t *tool) (id int, err error) {
	statement := "insert into tools (title, category, price, tool_owner) values ($1, $2, $3, $4) returning id"
	stmt, err := Db.Prepare(statement)
	fmt.Println(stmt)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(t.Title, t.Category, t.Price, t.ToolOwner).Scan(&t.ID)
	if err != nil {
		return
	}
	fmt.Println(t.ID)
	return t.ID, err
}

func rentToolUpdate(t *tool) {
	// Db.QueryRow(`update tools set available = false where id = $1`, id)
	Db.QueryRow(`update tools set start_date = $2, end_date = $3 where id = $1`, t.ID, t.StartDate, t.EndDate)
	// fmt.Println(t.EndDate)
	// fmt.Println(t.Title)
	// fmt.Println(t.ID)
	// fmt.Println(t.StartDate)
  // statement := "update tools set available = false, start_date = $2, end_date = $3 where id = $1"
	// stmt, err := Db.Prepare(statement)
	// fmt.Println(stmt)
	// if err != nil {
	// 	return
	// }
	// defer stmt.Close()
	// err = stmt.QueryRow(t.ID, t.StartDate, t.EndDate).Scan(&t.ID)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(t.ID)
	// return err
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

func selectTools() (toolOwners []toolOwner, err error) {
	rows, err := Db.Query(`select t.id as tool_id, t.title, t.category, t.price, t.tool_owner, 
		u.first_name, u.last_name, u.email, u.password_hash, u.phone_number,
		u.address_1, u.address_2, u.city, u.region, u.zipcode
		from tools t join users u on t.tool_owner = u.id`)
	if err != nil {
		return
	}
	for rows.Next() {
		to := toolOwner{}
		err = rows.Scan(&to.ID, &to.Title, &to.Category, &to.Price, &to.ToolOwner,
			&to.FirstName, &to.LastName, &to.Email, &to.PasswordHash, &to.PhoneNumber, 
			&to.Address1, &to.Address2,
			&to.City, &to.Region, &to.Zipcode)
		if err != nil {
			return
		}
		fmt.Println(to)
		toolOwners = append(toolOwners, to)
		fmt.Println(toolOwners)
	}
	rows.Close()
	return
}

func selectTool(id int) (to toolOwner, err error) {
	err = Db.QueryRow(`select t.id as tool_id, t.title, t.category, t.price, t.tool_owner, 
		u.first_name, u.last_name, u.email, u.password_hash, u.phone_number,
	  u.address_1, u.address_2, u.city, u.region, u.zipcode
		from tools t join users u on t.tool_owner = u.id where t.id = $1`, id).
		Scan(&to.ID, &to.Title, &to.Category, &to.Price, &to.ToolOwner, &to.FirstName,
		&to.LastName, &to.Email, &to.PasswordHash, &to.PhoneNumber, &to.Address1, &to.Address2,
		&to.City, &to.Region, &to.Zipcode)
	return
}

func selectUserTools(id int) (tools []tool, err error) {
	rows, err := Db.Query("select id, title, category, price, tool_owner from tools where tool_owner = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		t := tool{}
		err = rows.Scan(&t.ID, &t.Title, &t.Category, &t.Price, &t.ToolOwner)
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
	err = Db.QueryRow(`select id, title, category, price, tool_owner from tools where tool_owner = $1 and id = $2`, uID, tID).
		Scan(&t.ID, &t.Title, &t.Category, &t.Price, &t.ToolOwner)
	return
}