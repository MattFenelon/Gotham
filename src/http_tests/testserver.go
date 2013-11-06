package http_tests

import (
	"gotham/lib"
	"net/http/httptest"
	"persistence"
)

func startTestableApi() (server *httptest.Server, eventstore *persistence.InMemoryEventStore, filestore *persistence.InMemoryFileStore) {
	eventstore = persistence.NewInMemoryEventStore()
	filestore = persistence.NewInMemoryFileStore()
	exports := lib.Configure(eventstore, filestore)
	server = httptest.NewServer(exports.Handler)

	return server, eventstore, filestore
}
