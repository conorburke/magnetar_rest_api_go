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
