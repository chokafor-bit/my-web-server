package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserMessage struct {
	Message string `josn:message"`
	Status  int    `json"status"`
}

func HomePage(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-type", "application/json")

	data := UserMessage{
		Message: "welcome to simple http server in go, JSON data!",
		Status:  200,
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", HomePage)
	fmt.Println("starting application/json on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
