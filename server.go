package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ignaciocarvajal/goLangApiRest/connect"

	"github.com/gorilla/mux"
)

func main() {
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	log.Println("Server Up in port :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	user := connect.GetUser(userId)
	log.Println(user)
	json.NewEncoder(w).Encode(user)
}
