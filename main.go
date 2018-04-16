package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hit test route")
}

func users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := allUsers()
	fmt.Println("all users")
	fmt.Println(users)
	fmt.Println("second")
	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(js)
	// w.Write(js)
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

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
