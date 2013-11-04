package http_test

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"
	"testing"
)

func TestAddBook(t *testing.T) {
	t.Log("POST /books")

	server, store := startTestableApi()
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
	defer rsp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

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

type addBookRequest struct {
	mediaType          string
	body               string
	expectedStatusCode int
}

func TestAddInvalidBooks(t *testing.T) {
	requests := []addBookRequest{
		addBookRequest{mediaType: "invalid/mediatype", body: "", expectedStatusCode: 415},
		addBookRequest{mediaType: "", body: "", expectedStatusCode: 415},
		addBookRequest{mediaType: "multipart/", body: "", expectedStatusCode: 415},
		addBookRequest{mediaType: "multipart/form-data;", body: "", expectedStatusCode: 415},
		addBookRequest{mediaType: "multipart/form-data; boundary=", body: "", expectedStatusCode: 415},
		addBookRequest{mediaType: "multipart/form-data; boundary=abc", body: "", expectedStatusCode: 400},
		addBookRequest{mediaType: "multipart/form-data; boundary=abc", body: "--abc\r\nContent-Disposition: name=\"metadata\"\r\n", expectedStatusCode: 400},
		addBookRequest{mediaType: "multipart/form-data; boundary=abc", body: "--abc\r\nContent-Disposition: name=\"metadata\"\r\n\r\n\r\n--abc--", expectedStatusCode: 400},
		addBookRequest{mediaType: "multipart/form-data; boundary=abc", body: "--abc\r\nContent-Disposition: name=\"metadata\"\r\n\r\n                  \r\n--abc--", expectedStatusCode: 400},
		addBookRequest{mediaType: "multipart/form-data; boundary=abc", body: "--abc\r\nContent-Disposition: name=\"metadata\"\r\n\r\n{}\r\n--abc--", expectedStatusCode: 400},
		addBookRequest{mediaType: "multipart/form-data; boundary=abc", body: "--abc\r\nContent-Disposition: name=\"metadata\"\r\n\r\nthis-isnt-json\r\n--abc--", expectedStatusCode: 400},
	}

	for _, req := range requests {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Error("Panic occurred", err)
				}
			}()

			server, store := startTestableApi()
			defer server.Close()

			t.Logf("POST /books with mediatype: \"%v\", body: \"%v\"", req.mediaType, req.body)

			rsp, err := http.Post(server.URL+"/books", req.mediaType, strings.NewReader(req.body))
			defer rsp.Body.Close()
			if err != nil {
				t.Error(err)
				return
			}

			bodyBytes, _ := ioutil.ReadAll(rsp.Body)
			body := string(bodyBytes)

			t.Log("The response should be", req.expectedStatusCode, http.StatusText(req.expectedStatusCode))
			if rsp.StatusCode != req.expectedStatusCode {
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

			// TODO: add test to make sure the comic does not have any pages saved.
		}()
	}
}

func TestGetBooks(t *testing.T) {
	t.Log("GET /books")

	server, _ := startTestableApi()
	defer server.Close()

	rsp, err := http.Get(server.URL + "/books")
	defer rsp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

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
