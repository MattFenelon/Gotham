package lib

import (
	"domainservices"
	"gotham/lib/handlers"
	"net/http"
)

type Exports struct {
	Handler http.Handler
}

func Configure(store domainservices.EventStorer) (exports Exports) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", ServeHttp)
	serveMux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.BooksHandler(w, r, store)
	})

	exports = Exports{Handler: serveMux}
	return exports
}
