package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", isAuthorized(Home))
	log.Fatal(http.ListenAndServe(":8001", nil))
}
