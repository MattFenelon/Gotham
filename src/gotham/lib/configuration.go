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
	domain := domainservices.NewComicDomain(eventstore, filestore, viewstore)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", makeDomainHandleFunc(handlers.RootHandler, domain))
	serveMux.HandleFunc("/books", makeDomainHandleFunc(handlers.BooksHandler, domain))

	exports = Exports{Handler: serveMux}
	return exports
}

func makeDomainHandleFunc(f domainHandlerFunc, domain domainservices.ComicDomain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, domain)
	}
}

type domainHandlerFunc func(http.ResponseWriter, *http.Request, domainservices.ComicDomain)
