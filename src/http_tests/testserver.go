package http_tests

import (
	"gotham/lib"
	"net/http/httptest"
	"persistence"
)

func startTestableApi() (server *httptest.Server, es *persistence.InMemoryEventStore, fs *persistence.InMemoryFileStore, vs *persistence.InMemoryViewStore) {
	es = persistence.NewInMemoryEventStore()
	fs = persistence.NewInMemoryFileStore()
	vs = persistence.NewInMemoryViewStore()
	exports := lib.Configure(es, fs, vs)
	server = httptest.NewServer(exports.Handler)

	return server, es, fs, vs
}
