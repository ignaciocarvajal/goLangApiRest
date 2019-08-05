package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ignaciocarvajal/goLangApiRest/connect"

	"github.com/gorilla/mux"
	"github.com/ignaciocarvajal/goLangApiRest/structure"
)

func main() {
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user/new", newUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", updateUser).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", deleteUser).Methods("DELETE")

	log.Println("Server Up in port :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	status := "success"
	var message string
	user := connect.GetUser(userId)
	if user.ID <= 0 {
		status = "error"
		message = "user not found."
	}
	response := structure.Response{status, user, message}
	log.Println(user)
	json.NewEncoder(w).Encode(response)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	user := getUserRequest(r)
	response := structure.Response{"success", connect.CreateUser(user), ""}
	json.NewEncoder(w).Encode(response)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	userUpdate := getUserRequest(r)
	response := structure.Response{"success", connect.UpdateUser(userId, userUpdate), ""}
	json.NewEncoder(w).Encode(response)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	var user = structure.User{}
	connect.DeleteUser(userId)
	response := structure.Response{"success", user, ""}
	json.NewEncoder(w).Encode(response)

}

func getUserRequest(r *http.Request) structure.User {
	var user = structure.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
