package main

import (
	"net/http"
	hello "read-and-stockwords/Hello"
	"read-and-stockwords/db"

	"github.com/gorilla/mux"
)

func main() {
	db.CreateDb()
	r := mux.NewRouter()
	r.HandleFunc("/", hello.Hello)
	http.ListenAndServe(":8080", r)
}
