package http_tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"gotham/lib"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"persistence"
	"testing"
)

func Test404(t *testing.T) {
	exports := lib.Configure(persistence.NewInMemoryEventStore())
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	response, err := http.Get(server.URL + "/rubbish")
	if err != nil {
		t.Fatalf("Error on HTTP GET to %v: %v", server.URL, err)
	}

	Convey("GET /rubbish", t, func() {
		Convey("The response should be 404 Not Found", func() {
			So(response.StatusCode, ShouldEqual, 404)
		})
		Convey("The Content-Type should be text/plain; charset=utf-8", func() {
			So(response.Header.Get("Content-Type"), ShouldEqual, "text/plain; charset=utf-8")
		})
		Convey("The response body should state that the resource could not be found", func() {
			defer response.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			body := string(bodyBytes)
			So(body, ShouldEqual, "404 page not found\n")
		})
	})
}
