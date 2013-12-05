package http_tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAddBook(t *testing.T) {
	t.Log("POST /books")

	api := newTestableApi()
	defer api.Close()

	metadata := `{
	"seriesTitle": "Prophet",
	"title": "Prophet 31",
	"publishedDate": "2012-11-28T00:00:00Z",
	"writtenBy": ["Brandon Graham", "Simon Roy", "Giannis Milonogiannis"],
	"artBy": ["Giannis Milonogiannis"],
	"blurb": "Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance."
}`

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writeMetadata(writer, metadata)
	image1 := writeImageData(t, writer, "testdata\\0.jpg")
	image2 := writeImageData(t, writer, "testdata\\1.jpg")
	writer.Close()

	rsp, err := http.Post(api.URL()+"/books", writer.FormDataContentType(), &buffer)
	defer rsp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	body := string(bodyBytes)

	actualEvents := api.es.GetAllEvents()

	t.Log("The response should be 204 No Content")
	if rsp.StatusCode != 204 {
		t.Errorf("\tExpected 204 but was %v", rsp.StatusCode)
	}

	t.Log("The response body should be empty")
	if body != "" {
		t.Errorf("\tExpected \"\" but was %v", body)
	}

	t.Log("The comic should be persisted")
	if len(actualEvents) != 1 {
		t.Errorf("\tExpected 1 items but contained %v", len(actualEvents))
	}

	t.Log("The comic's title should be persisted")
	if actualEvents[0].Title != "Prophet 31" {
		t.Errorf("\tExpected %v but contained %v", "Prophet 31", actualEvents[0].Title)
	}

	t.Log("The comic's series title should be persisted")
	if actualEvents[0].SeriesTitle != "Prophet" {
		t.Errorf("\tExpected %v but contained %v", "Prophet", actualEvents[0].SeriesTitle)
	}

	t.Log("The comic's published date should be persisted")
	expectedDate := time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC)
	if actualEvents[0].PublishedDate != expectedDate {
		t.Errorf("\tExpected %v but contained %v", expectedDate, actualEvents[0].PublishedDate)
	}

	t.Log("The comic's writers should be persisted")
	expectedWriters := []string{"Brandon Graham", "Simon Roy", "Giannis Milonogiannis"}
	if reflect.DeepEqual(expectedWriters, actualEvents[0].WrittenBy) == false {
		t.Errorf("\tExpected %v but contained %v", expectedWriters, actualEvents[0].WrittenBy)
	}

	t.Log("The comic's artists should be persisted")
	expectedArtists := []string{"Giannis Milonogiannis"}
	if reflect.DeepEqual(expectedArtists, actualEvents[0].ArtBy) == false {
		t.Errorf("\tExpected %v but contained %v", expectedArtists, actualEvents[0].ArtBy)
	}

	t.Log("The comic's blurb should be persisted")
	expectedBlurb := "Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance."
	if actualEvents[0].Blurb != expectedBlurb {
		t.Errorf("\tExpected %v but contained %v", expectedBlurb, actualEvents[0].Blurb)
	}

	t.Log("The comic pages should be persisted using the filenames specified in the form")
	actualFiles := getAllFilenames(api.fs)
	expectedFiles := []string{"0.jpg", "1.jpg"}
	if reflect.DeepEqual(actualFiles, expectedFiles) == false {
		t.Errorf("\tExpected %v but contained %v", expectedFiles, actualFiles)
	}

	t.Log("The contents of the persisted comic pages should match the uploaded images")
	actualContents := getAllFileContents(api.fs)
	expectedContents := [][]byte{image1, image2}
	if reflect.DeepEqual(actualContents, expectedContents) == false {
		t.Errorf("\tExpected %v but was %v", len(actualContents), len(expectedContents))
	}
}

func writeMetadata(w *multipart.Writer, metadata string) {
	metadataHeader := make(textproto.MIMEHeader)
	metadataHeader.Set("Content-Disposition", "form-data; name=\"metadata\"")
	metadataHeader.Set("Content-Type", "application/json")
	metadataPart, _ := w.CreatePart(metadataHeader)
	metadataPart.Write([]byte(metadata))
}

func writeImageData(t *testing.T, w *multipart.Writer, imagepath string) []byte {
	_, filename := filepath.Split(imagepath)
	pageImageHeader := make(textproto.MIMEHeader)
	pageImageHeader.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"page\"; filename=\"%v\"", filename))
	pageImageHeader.Set("Content-Type", "image/jpeg")

	abs, _ := filepath.Abs(imagepath)
	contents, err := ioutil.ReadFile(abs)
	if err != nil {
		t.Fatal(err)
	}

	imagePart, _ := w.CreatePart(pageImageHeader)
	imagePart.Write(contents)
	return contents
}

func getAllFilenames(filestore *testFileStore) []string {
	keys, _ := filestore.GetAllKeys()
	files := []string{}
	for _, key := range keys {
		filenames, _ := filestore.GetFilenames(key)
		files = append(files, filenames...)
	}

	return files
}

func getAllFileContents(filestore *testFileStore) [][]byte {
	keys, _ := filestore.GetAllKeys()
	files := [][]byte{}
	for _, key := range keys {
		filenames, _ := filestore.GetFilenames(key)

		for _, filename := range filenames {
			f, _ := filestore.Open(key + "/" + filename)
			defer f.Close()
			content, _ := ioutil.ReadAll(f)

			files = append(files, content)
		}
	}

	return files
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

			api := newTestableApi()
			defer api.Close()

			t.Logf("POST /books with mediatype: \"%v\", body: \"%v\"", req.mediaType, req.body)

			rsp, err := http.Post(api.URL()+"/books", req.mediaType, strings.NewReader(req.body))
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

			t.Log("The comic data should not be persisted")
			actualEvents := api.es.GetAllEvents()
			if len(actualEvents) != 0 {
				t.Errorf("\tExpected 0 items but contained %v", actualEvents)
			}

			// TODO: add test to make sure the comic does not have any pages saved.
		}()
	}
}

func TestGetBooks(t *testing.T) {
	t.Log("GET /books")

	api := newTestableApi()
	defer api.Close()

	rsp, err := http.Get(api.URL() + "/books")
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
