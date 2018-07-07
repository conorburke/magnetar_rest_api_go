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
	// 	FirstName: "Jim",
	// 	LastName: "Lahey",
	// 	Email: "theliqour@gmail.com",
	// 	PasswordHash: "password",
	// 	PhoneNumber:"123456789",
	// 	Age: 58,
	// 	Address1: "27 Boneventure",
	// 	Address2: "",
	// 	City: "Sunnyvale",
	// 	Region: "Ontario",
	// 	Zipcode: 12345,
	// }

	// fmt.Println(u)
	// u.createUser()
	// fmt.Println(u)

	mux := httprouter.New()
	mux.GET("/api/test", test)
	mux.GET("/api/users", getUsers)
	mux.GET("/api/users/:id", getUser)
	mux.GET("/api/tools", getTools)
	mux.GET("/api/tools/:id", getTool)
	mux.GET("/api/users/:id/tools", getUserTools)
	//TODO: need to figure out wildcard conflict, want to use users instead of user below
	mux.GET("/api/user/:userID/tools/:toolID", getUserTool)
	mux.DELETE("/api/tools/:id", deleteTool)
	mux.POST("/api/users", createUser)
	mux.POST("/api/tools", createTool)
	mux.POST("/api/rent/:id", rentTool)

	server := http.Server{
		Addr:    "127.0.0.1:" + port,
		Handler: mux,
	}

	server.ListenAndServe()

}
