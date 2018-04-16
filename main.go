package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	port := "8080"
	fmt.Printf("server up and running on %s \n", port)

	// u := user{
	// 	firstName: "Randy",
	// 	lastName: "Lahey",
	// 	email: "gutactular@gmail.com",
	// 	passwordHash: "password",
	// 	phoneNumber:"123456789",
	// 	age: 58,
	// 	address1: "27 Boneventure",
	// 	address2: "",
	// 	city: "Sunnyvale",
	// 	region: "Ontario",
	// 	zipcode: 12345,
	// }

	// fmt.Println(u)
	// u.createUser()
	// fmt.Println(u)

	mux := httprouter.New()
	mux.GET("/test", test)
	mux.GET("/users", users)

	server := http.Server{
		Addr:    "127.0.0.1:" + port,
		Handler: mux,
	}

	server.ListenAndServe()

}
