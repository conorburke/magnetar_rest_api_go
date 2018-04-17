package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strconv"
)

func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hit test route")
}

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := selectUsers()
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		panic(err)
	}
	u, err := selectUser(userID)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getTools(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tools, err := selectTools()
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(tools)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getTool(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	toolID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		panic(err)
	}
	t, err := selectUser(toolID)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getUserTools(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID, err := strconv.Atoi(p.ByName("id"))
	tools, err := selectUserTools(userID)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(tools)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getUserTool(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID, err := strconv.Atoi(p.ByName("userID"))
	if err != nil {
		panic(err)
	}
	toolID, err := strconv.Atoi(p.ByName("toolID"))
	if err != nil {
		panic(err)
	}
	t, err := selectUserTool(userID, toolID)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
