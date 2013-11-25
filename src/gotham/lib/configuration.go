package lib

import (
	"domain"
	"gotham/lib/handlers"
	"net/http"
)

type Exports struct {
	Handler http.Handler
}

func Configure(eventstore domain.EventStorer, filestore FileStore, viewstore domain.ViewGetStorer) (exports Exports) {
	domain := domain.NewComicDomain(eventstore, filestore, viewstore)

	// TODO: The API is vulnerable to Host header spoofing because the Host header is used in
	// returned links. The vulnerability can be closed off by ensuring that the API only responds
	// to whitelisted hosts.
	// Tests don't use the same host as the real API. Whatever solution is used to plug this hole it needs to
	// not break the tests.

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", makeDomainHandleFunc(handlers.RootHandler, domain))
	serveMux.HandleFunc("/books", makeDomainHandleFunc(handlers.BooksHandler, domain))
	serveMux.HandleFunc("/books/", makeDomainHandleFunc(handlers.BookHandler, domain))
	serveMux.Handle("/pages/", makeFilestoreHandler("/pages/", filestore))

	exports = Exports{Handler: serveMux}
	return exports
}

func makeDomainHandleFunc(f domainHandleFunc, domain domain.ComicDomain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, domain)
	}
}

type domainHandleFunc func(http.ResponseWriter, *http.Request, domain.ComicDomain)
