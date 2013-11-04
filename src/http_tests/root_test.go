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

func TestGetRootResource(t *testing.T) {
	exports := lib.Configure(persistence.NewInMemoryEventStore())
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Error on HTTP GET to %v: %v", server.URL, err)
	}

	Convey("GET /", t, func() {
		Convey("The response should be 200 OK", func() {
			So(response.StatusCode, ShouldEqual, 200)
		})

		Convey("The Content-Type should be application/json", func() {
			So(response.Header.Get("Content-Type"), ShouldEqual, "application/json")
		})

		Convey("The response body should include all comics in JSON format", func() {
			defer response.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			body := string(bodyBytes)
			So(body, ShouldEqual, `{
	"seriesset": [
		{
			"title": "Prophet",
			"links": {
				"via": {"href": "/series/1"},
				"{docHost}/rel/seriesimage": {"href": "/images/1.jpg"}
			},
		},
		{
			"title": "Jupiter's Legacy",
			"links": {
				"via": {"href": "/series/2"},
				"{docHost}/rel/seriesimage": {"href": "/images/2.jpg"}
			},
		}
	]
}`)
		})
	})
}
