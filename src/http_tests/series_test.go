package http_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGetSeries(t *testing.T) {
	t.Log("When multiple comics of the same series are added")

	api := newTestableApi()
	defer api.Close()
	comics := domain.NewComicDomain(api.es, api.fs, api.vs)

	expectedBlurb1 := `Nicolas Lash is in the deepest trouble possible, and there's only one person who can save him now... the problem is, trouble is her business.`
	expectedBlurb2 := `Rock 'n' roll, robbery, and ritual murder collide in '90s Seattle as the fourth arc of FATALE comes to a shattering conclusion... and the secret identity of Nic's present day helper is revealed. Is he who he claims to be?`
	expectedDate1 := time.Date(2014, time.February, 5, 0, 0, 0, 0, time.UTC)
	expectedDate2 := time.Date(2014, time.January, 8, 0, 0, 0, 0, time.UTC)
	expectedWriters := []string{"Ed Brubaker"}
	expectedArtists := []string{"Sean Philips", "Elizabeth Breitweiser"}
	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()
	comics.AddComic(id2, "Fatale", "Fatale 19", []string{"1.jpg"}, []string{"source//path//1.jpg"}, expectedWriters, expectedArtists, expectedDate2, expectedBlurb2)
	comics.AddComic(id1, "Fatale", "Fatale 20", []string{"0.jpg"}, []string{"source//path//0.jpg"}, expectedWriters, expectedArtists, expectedDate1, expectedBlurb1)

	seriesUri := api.URL() + "/series/Fatale"

	t.Logf("\tGET %v", seriesUri)
	rsp, err := http.Get(seriesUri)
	if err != nil {
		t.Fatalf("\t\tErr on GET to %v: %v", seriesUri, err)
	}
	defer rsp.Body.Close()
	raw, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actualBody := string(raw)

	t.Log("The response should be 200 OK")
	if rsp.StatusCode != http.StatusOK {
		t.Errorf("\tThe response status code should be %v but was %v", http.StatusOK, rsp.StatusCode)
	}

	t.Log("The response body should contain the series' comics in LIFO order")
	expectedBody := `{"title":"Fatale",` +
		`"books":[` +
		`{` +
		`"title":"Fatale 20","publishedDate":"2014-02-05T00:00:00Z","writtenBy":["Ed Brubaker"],"artBy":["Sean Philips","Elizabeth Breitweiser"],` +
		`"blurb":"Nicolas Lash is in the deepest trouble possible, and there's only one person who can save him now... the problem is, trouble is her business.",` +
		`"links":[{"rel":"self","href":"` + api.URL() + `/books/` + id1.String() + `"},{"rel":"image","href":"` + api.URL() + `/pages/` + id1.String() + `/0.jpg"}]` +
		`},` +
		`{` +
		`"title":"Fatale 19","publishedDate":"2014-01-08T00:00:00Z","writtenBy":["Ed Brubaker"],"artBy":["Sean Philips","Elizabeth Breitweiser"],` +
		`"blurb":"Rock 'n' roll, robbery, and ritual murder collide in '90s Seattle as the fourth arc of FATALE comes to a shattering conclusion... and the secret identity of Nic's present day helper is revealed. Is he who he claims to be?",` +
		`"links":[{"rel":"self","href":"` + api.URL() + `/books/` + id2.String() + `"},{"rel":"image","href":"` + api.URL() + `/pages/` + id2.String() + `/1.jpg"}]` +
		`}` +
		`]}` + "\n"
	if expectedBody != actualBody {
		t.Errorf("Expected %v but was %v", expectedBody, actualBody)
	}
}

func TestGetSeriesWithEncodedSpaces(t *testing.T) {
	t.Log("When a comic is added")

	api := newTestableApi()
	defer api.Close()
	comics := domain.NewComicDomain(api.es, api.fs, api.vs)

	expectedDate := time.Date(2014, time.January, 8, 0, 0, 0, 0, time.UTC)
	id := uuid.NewRandom()
	comics.AddComic(id, "The Walking Dead", "The Walking Dead 1", []string{"0.jpg"}, []string{"source//path//0.jpg"}, []string{}, []string{}, expectedDate, "")

	seriesUri := api.URL() + "/series/The%20Walking%20Dead"

	t.Logf("\tGET %v", seriesUri)
	rsp, err := http.Get(seriesUri)
	if err != nil {
		t.Fatalf("\t\tErr on GET to %v: %v", seriesUri, err)
	}
	defer rsp.Body.Close()
	raw, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actualBody := string(raw)

	t.Log("The response should be 200 OK")
	if rsp.StatusCode != http.StatusOK {
		t.Errorf("\tThe response status code should be %v but was %v", http.StatusOK, rsp.StatusCode)
	}

	t.Log("The response body should contain the series' data")
	expectedBody := `{"title":"The Walking Dead",` +
		`"books":[` +
		`{` +
		`"title":"The Walking Dead 1","publishedDate":"2014-01-08T00:00:00Z","writtenBy":[],"artBy":[],` +
		`"blurb":"",` +
		`"links":[{"rel":"self","href":"` + api.URL() + `/books/` + id.String() + `"},{"rel":"image","href":"` + api.URL() + `/pages/` + id.String() + `/0.jpg"}]` +
		`}` +
		`]}` + "\n"
	if expectedBody != actualBody {
		t.Errorf("Expected %v but was %v", expectedBody, actualBody)
	}
}

func TestGetSeries404(t *testing.T) {
	t.Log("When retrieving an unknown series")

	api := newTestableApi()
	defer api.Close()

	rsp, err := http.Get(api.URL() + "/series/ThisIsNotASeries")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("The response should be 404 Not Found")
	if rsp.StatusCode != http.StatusNotFound {
		t.Errorf("\tThe response status code should be %v but was %v", http.StatusNotFound, rsp.StatusCode)
	}
}
