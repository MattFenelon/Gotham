package main

import (
	"io"
	"log"
	"net/http"
)

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	response :=
		`{
	"series" : [
		{
			"title" : "Prophet"
		},
		{
			"title" : "Jupiter's Legacy"
		}
	]
}`
	w.Header().Add("Content-Type", "application/json")

	io.WriteString(w, response)
}

func main() {
	http.HandleFunc("/", ServeHttp)
	err := http.ListenAndServe(":7001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
