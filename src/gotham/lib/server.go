package lib

import (
	"log"
	"net/http"
	"persistence/filestore"
	"persistence/riak"
)

func StartServer() {
	riakCluster := []string{"127.0.0.1:8087"}
	riakClientId := "gotham"

	log.Println("Gotham is starting...")

	eventstore := riak.NewRiakEventStore(riakCluster, riakClientId)
	viewstore := riak.NewViewStore(riakCluster, riakClientId)
	filestore := filestore.NewLocalFileStore("c:\\gothamfs")
	exports := Configure(eventstore, filestore, viewstore)

	err := http.ListenAndServe(":7001", exports.Handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
