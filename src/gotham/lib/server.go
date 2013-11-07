package lib

import (
	"log"
	"net/http"
	"persistence"
	"persistence/filestore"
	"persistence/riak"
)

func StartServer() {
	log.Println("Gotham is starting...")

	eventstore := riak.NewRiakEventStore([]string{"127.0.0.1:8087"}, "httpapi")
	filestore := filestore.NewLocalFileStore("c:\\gothamfs")
	viewstore := persistence.NewInMemoryViewStore()
	exports := Configure(eventstore, filestore, viewstore)

	err := http.ListenAndServe(":7001", exports.Handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
