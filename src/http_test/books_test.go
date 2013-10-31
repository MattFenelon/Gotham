package http_test

import (
	"gotham/lib"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"persistence"
	"strings"
	"testing"
)

func TestAddBook(t *testing.T) {
	t.Log("POST /books/")
	t.Skip("Not implemented")
	store := persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	reader := strings.NewReader(`{
	"seriesTitle": "Prophet"
	"title": "Prophet 31"
}`)
	rsp, err := http.Post(server.URL+"/books", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}

	defer rsp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	body := string(bodyBytes)

	t.Log("The response should be 204 No Content")
	if rsp.StatusCode != 204 {
		t.Errorf("\tExpected 204 but was %v", rsp.StatusCode)
	}

	t.Log("The response body should be empty")
	if body != "" {
		t.Errorf("\tExpected \"\" but was %v", body)
	}

	t.Log("The comic should be persisted")
	actualEvents := store.GetAllEvents()
	if len(actualEvents) == 0 || actualEvents[0].Title != "Prophet 31" {
		t.Errorf("\tExpected 1 items but was %v", actualEvents)
	}
}
