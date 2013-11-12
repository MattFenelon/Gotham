package http_tests

import (
	"gotham/lib"
	"io/ioutil"
	"net/http/httptest"
	"persistence"
)

type api struct {
	es     *persistence.InMemoryEventStore
	fs     *testFileStore
	vs     *persistence.InMemoryViewStore
	server *httptest.Server
}

func newTestableApi() *api {
	path, _ := ioutil.TempDir("", "httpapitests_filestore")

	es := persistence.NewInMemoryEventStore()
	fs := newTestFileStore(path)
	vs := persistence.NewInMemoryViewStore()
	exports := lib.Configure(es, fs, vs)
	server := httptest.NewServer(exports.Handler)

	return &api{
		es,
		fs,
		vs,
		server,
	}
}

func (api *api) URL() string {
	return api.server.URL
}

func (api *api) Close() {
	api.server.Close()
	api.fs.close()
}
