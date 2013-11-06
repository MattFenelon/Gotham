package main

import (
	"gotham/lib"
	"log"
	"net/http"
	"persistence/filestore"
	"persistence/riak"
)

func main() {
	log.Println("Gotham is starting...")

	eventstore := riak.NewRiakEventStore([]string{"127.0.0.1:8087"}, "httpapi")
	filestore := filestore.NewLocalFileStore("c:\\gothamfs")
	exports := lib.Configure(eventstore, filestore)

	err := http.ListenAndServe(":7001", exports.Handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
