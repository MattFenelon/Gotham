package http_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBookGetBook(t *testing.T) {
	t.Log("When a comic has been added")

	api := newTestableApi()
	defer api.Close()

	comics := domain.NewComicDomain(api.es, api.fs, api.vs)
	fataleId := uuid.NewRandom()
	comics.AddComic(fataleId, "Fatale", "Fatale 18", []string{"0.jpg", "1.jpg"}, []string{"testdata\\0.jpg", "testdata\\1.jpg"})

	bookUri := api.URL() + "/books/" + fataleId.String()

	t.Logf("\tGET %v", bookUri)
	rsp, err := http.Get(bookUri)
	if err != nil {
		t.Errorf("\t\tErr on GET to %v: %v", bookUri, err)
	}
	defer rsp.Body.Close()
	raw, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actualBody := string(raw)

	t.Log("\t\tThe response should be 200 OK")
	if rsp.StatusCode != 200 {
		t.Error("\t\t\tExpected 200 but was", rsp.StatusCode)
	}

	t.Log("\t\tThe response body should contain to the comic's page images")
	expectedBody := `{"links":[` +
		`{"rel":"item","href":"` + api.URL() + `/pages/` + fataleId.String() + `/0.jpg"},` +
		`{"rel":"item","href":"` + api.URL() + `/pages/` + fataleId.String() + `/1.jpg"}]}` + "\n"

	if expectedBody != actualBody {
		t.Errorf("Expected %v but was %v", expectedBody, actualBody)
	}
}

func TestBook404(t *testing.T) {
	t.Log("When retrieving an unknown comic")

	api := newTestableApi()
	defer api.Close()

	rsp, err := http.Get(api.URL() + "/books/rubbish")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("The status code should be 404 Not Found")
	if rsp.StatusCode != http.StatusNotFound {
		t.Errorf("\tExpected %v but was %v", 404, rsp.StatusCode)
	}
}
