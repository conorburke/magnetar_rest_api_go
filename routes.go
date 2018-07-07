package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
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

func createTool(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(body)

	var t tool

	err = json.Unmarshal(body, &t)

	if err != nil {
		panic(err)
	}

	id, err := insertTool(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func rentTool(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
	// toolID, err := strconv.Atoi(p.ByName("id"))
	// if err != nil {
	// 	panic(err)
	// }

	// rentToolUpdate(toolID)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(body)

	var t tool

	err = json.Unmarshal(body, &t)

	if err != nil {
		panic(err)
	}

	// err = rentToolUpdate(&t)

	// if err != nil {
	// 	panic(err)
	// }

	rentToolUpdate(&t)
}


func createUser(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(body)

	var u user

	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	log.Println(u.FirstName)

	id, err := insertUser(&u)

	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	user, err := selectUser(id)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func deleteTool(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	toolID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		panic(err)
	}
	destroyTool(toolID)
}

func getTools(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	toolOwners, err := selectTools()
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(toolOwners)
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
	toolOwner, err := selectTool(toolID)

	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(toolOwner)
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
