package lib

import (
	"domainservices"
	"gotham/lib/handlers"
	"net/http"
)

type Exports struct {
	Handler http.Handler
}

func Configure(eventstore domainservices.EventStorer, filestore domainservices.FileStorer, viewstore domainservices.ViewGetStorer) (exports Exports) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handlers.RootHandler)
	serveMux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.BooksHandler(w, r, eventstore, filestore, viewstore)
	})

	exports = Exports{Handler: serveMux}
	return exports
}
