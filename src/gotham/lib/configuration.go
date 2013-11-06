package lib

import (
	"domainservices"
	"gotham/lib/handlers"
	"net/http"
)

type Exports struct {
	Handler http.Handler
}

func Configure(eventstore domainservices.EventStorer, filestore domainservices.FileStorer) (exports Exports) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", ServeHttp)
	serveMux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.BooksHandler(w, r, eventstore, filestore)
	})

	exports = Exports{Handler: serveMux}
	return exports
}
