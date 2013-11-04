package http_tests

import (
	"gotham/lib"
	"net/http/httptest"
	"persistence"
)

func startTestableApi() (server *httptest.Server, store *persistence.InMemoryEventStore) {
	store = persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server = httptest.NewServer(exports.Handler)

	return server, store
}
