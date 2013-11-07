package http_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetRootResource(t *testing.T) {
	t.Log("When the root resource contains comics")

	server, eventstore, filestore, viewstore := startTestableApi()
	defer server.Close()

	comics := domainservices.NewComicDomain(eventstore, filestore, viewstore)
	fataleId := uuid.NewRandom()
	walkingDeadId := uuid.NewRandom()
	comics.AddComic(fataleId, "Fatale", "Fatale 18", []string{"0.jpg", "1.jpg"}, []string{"testdata\\0.jpg", "testdata\\1.jpg"})
	comics.AddComic(walkingDeadId, "The Walking Dead", "The Walking Dead 115", []string{"0.jpg", "1.jpg"}, []string{"testdata\\0.jpg", "testdata\\1.jpg"})

	t.Log("\tGET /")

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Error on HTTP GET to %v: %v", server.URL, err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	actualBody := string(bodyBytes)

	t.Log("\tThe response should be 200 OK")
	if response.StatusCode != 200 {
		t.Error("\tExpected 200 but was", response.StatusCode)
	}

	t.Log("\tThe Content-Type should be application/json")
	if response.Header.Get("Content-Type") != "application/json" {
		t.Error("\tExpected application/json but was", response.Header.Get("Content-Type"))
	}

	t.Log("\tThe response body should include all comics in JSON format")
	expectedBody := `{
	"series": [
		{
			"title": "Fatale",
			"links": {
				"seriesimage": {"href": "http://gotham/pages/` + fataleId.String() + `/0.jpg"}
			}
		},
		{
			"title": "The Walking Dead",
			"links": {
				"seriesimage": {"href": "http://gotham/pages/` + walkingDeadId.String() + `/0.jpg"}
			}
		}
	]
}`
	if actualBody != expectedBody {
		t.Errorf("\tExpected %v but was %v", expectedBody, actualBody)
	}
}
