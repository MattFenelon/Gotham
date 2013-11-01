package main

import (
	"gotham/lib"
	"log"
	"net/http"
	"persistence/riak"
)

func main() {
	exports := lib.Configure(riak.NewRiakEventStore([]string{"127.0.0.1:8087"}, "httpapi"))

	err := http.ListenAndServe(":7001", exports.Handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
