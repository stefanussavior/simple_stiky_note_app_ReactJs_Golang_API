package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	var user Users
	var arr_user []Users

	db := connect()
	defer db.Close()

	rows, err := db.Query("select id, first_name, last_name from person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, user)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arr_user)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getProducts", returnAllProducts).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
