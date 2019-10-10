package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	home := http.HandlerFunc(Home)
	http.Handle("/home", isAuthorized(home))
	log.Fatal(http.ListenAndServe(":8001", nil))
}
