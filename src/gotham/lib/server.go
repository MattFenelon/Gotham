package lib

import (
	"log"
	"net/http"
	"persistence/filestore"
	"persistence/riak"
)

func StartServer() {
	cluster := []string{"127.0.0.1:8087"}
	clientId := "gotham"

	log.Println("Gotham is starting...")

	eventstore := riak.NewRiakEventStore(cluster, clientId)
	viewstore := riak.NewViewStore(cluster, clientId)
	filestore := filestore.NewLocalFileStore("c:\\gothamfs")
	exports := Configure(eventstore, filestore, viewstore)

	err := http.ListenAndServe(":7001", exports.Handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
