package http_test

import (
	"bytes"
	"gotham/lib"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"path/filepath"
	"persistence"
	"strings"
	"testing"
)

func TestAddBook(t *testing.T) {
	t.Log("POST /books")

	store := persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	metadata := `{
	"seriesTitle": "Prophet",
	"title": "Prophet 31"
}`

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writeMetadata(writer, metadata)
	writeImageData(t, writer, "testdata\\Prophet 30 Cover Image.jpg")
	writer.Close()

	rsp, err := http.Post(server.URL+"/books", writer.FormDataContentType(), &buffer)
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
		t.Errorf("\tExpected 1 items but contained %v", actualEvents)
	}

	t.Log("The comic pages should be persisted")
	t.Error("\tExpected 1 items but contained 0")
}

func TestAddBookWithInvalidContentType(t *testing.T) {
	t.Log("POST /books")

	store := persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	reader := strings.NewReader(`{
	"seriesTitle": "Prophet",
	"title": "Prophet 31"
}`)
	rsp, err := http.Post(server.URL+"/books", "invalid/mediatype", reader)
	if err != nil {
		t.Fatal(err)
	}

	defer rsp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	body := string(bodyBytes)

	t.Log("The response should be 415 Unsupported Media Type")
	if rsp.StatusCode != 415 {
		t.Errorf("\tExpected 415 but was %v", rsp.StatusCode)
	}

	t.Log("The response body should be empty")
	if body != "" {
		t.Errorf("\tExpected \"\" but was %v", body)
	}

	t.Log("The comic should not be persisted")
	actualEvents := store.GetAllEvents()
	if len(actualEvents) != 0 {
		t.Errorf("\tExpected 0 items but contained %v", actualEvents)
	}
}

func TestAddBookWithInvalidJSON(t *testing.T) {
	t.Log("POST /books")

	store := persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	reader := strings.NewReader(`complete-and-utter-rubbish`)
	rsp, err := http.Post(server.URL+"/books", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}

	defer rsp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	body := string(bodyBytes)

	t.Log("The response should be 400 Bad Request")
	if rsp.StatusCode != 400 {
		t.Errorf("\tExpected 400 but was %v", rsp.StatusCode)
	}

	t.Log("The response body should be empty")
	if body != "" {
		t.Errorf("\tExpected \"\" but was %v", body)
	}

	t.Log("The comic should not be persisted")
	actualEvents := store.GetAllEvents()
	if len(actualEvents) != 0 {
		t.Errorf("\tExpected 0 items but contained %v", actualEvents)
	}
}

func TestAddBookWithEmptyJSON(t *testing.T) {
	t.Log("POST /books")

	store := persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	reader := strings.NewReader(`{}`)
	rsp, err := http.Post(server.URL+"/books", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}

	defer rsp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	body := string(bodyBytes)

	t.Log("The response should be 400 Bad Request")
	if rsp.StatusCode != 400 {
		t.Errorf("\tExpected 400 but was %v", rsp.StatusCode)
	}

	t.Log("The response body should be empty")
	if body != "" {
		t.Errorf("\tExpected \"\" but was %v", body)
	}

	t.Log("The comic should not be persisted")
	actualEvents := store.GetAllEvents()
	if len(actualEvents) != 0 {
		t.Errorf("\tExpected 0 items but contained %v", actualEvents)
	}
}

func TestGet(t *testing.T) {
	t.Log("GET /books")

	store := persistence.NewInMemoryEventStore()
	exports := lib.Configure(store)
	server := httptest.NewServer(exports.Handler)
	defer server.Close()

	rsp, err := http.Get(server.URL + "/books")
	if err != nil {
		t.Fatal(err)
	}

	defer rsp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	body := string(bodyBytes)

	t.Log("The response should be 405 Method Not Allowed")
	if rsp.StatusCode != 405 {
		t.Errorf("\tExpected 405 but was %v", rsp.StatusCode)
	}

	t.Log("The response body should be empty")
	if body != "" {
		t.Errorf("\tExpected \"\" but was %v", body)
	}
}

func writeMetadata(w *multipart.Writer, metadata string) {
	metadataHeader := make(textproto.MIMEHeader)
	metadataHeader.Set("Content-Disposition", "form-data; name=\"metadata\"")
	metadataHeader.Set("Content-Type", "application/json")
	metadataPart, _ := w.CreatePart(metadataHeader)
	metadataPart.Write([]byte(metadata))
}

func writeImageData(t *testing.T, w *multipart.Writer, imagepath string) {
	pageImageHeader := make(textproto.MIMEHeader)
	pageImageHeader.Set("Content-Disposition", "form-data; name=\"page\"")
	pageImageHeader.Set("Content-Type", "image/jpeg")

	abs, _ := filepath.Abs(imagepath)
	contents, err := ioutil.ReadFile(abs)
	if err != nil {
		t.Fatal(err)
	}

	imagePart, _ := w.CreatePart(pageImageHeader)
	imagePart.Write(contents)
}
