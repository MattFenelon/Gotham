package main

import (
	"gotham/lib"
	"log"
	"net/http"
)

func main() {
	exports := lib.Configure(nil)

	err := http.ListenAndServe(":7001", exports.Handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
