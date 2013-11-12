package http_tests

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRootGet(t *testing.T) {
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
	expectedBody :=
		`{"series":[` +
			`{"title":"The Walking Dead","links":{"seriesimage":{"href":"http://gotham/pages/` + walkingDeadId.String() + `/0.jpg"}}},` +
			`{"title":"Fatale","links":{"seriesimage":{"href":"http://gotham/pages/` + fataleId.String() + `/0.jpg"}}}` +
			`]}` + "\n"

	if actualBody != expectedBody {
		t.Errorf("\tExpected %v but was %v", expectedBody, actualBody)
	}
}

type root struct {
	Series []struct {
		Links struct {
			Seriesimage struct {
				Href string
			}
		}
	}
}

func TestRootGetSeriesImage(t *testing.T) {
	t.Log("When the root resource contains comics with different series images")

	server, eventstore, filestore, viewstore := startTestableApi()
	defer server.Close()

	comics := domainservices.NewComicDomain(eventstore, filestore, viewstore)
	fataleId := uuid.NewRandom()
	walkingDeadId := uuid.NewRandom()
	comics.AddComic(fataleId, "Fatale", "Fatale 18", []string{"0.jpg"}, []string{"testdata\\0.jpg"})
	comics.AddComic(walkingDeadId, "The Walking Dead", "The Walking Dead 115", []string{"1.jpg"}, []string{"testdata\\1.jpg"})

	expectedImages, err := getFileContents("testdata\\1.jpg", "testdata\\0.jpg") // The order matters. Series have LIFO ordering
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("\tGET %v", server.URL)
	root := getRoot(t, server.URL)
	if len(root.Series) != 2 {
		t.Fatalf("\tExpected 2 series but was %v", len(root.Series))
	}

	for i, s := range root.Series {
		expectedImage := expectedImages[i]
		checkSeriesImage(t, s.Links.Seriesimage.Href, expectedImage)
	}
}

func checkSeriesImage(t *testing.T, imageUri string, expectedImage []byte) {
	t.Logf("\t\tGET %v", imageUri)

	rsp, err := http.Get(imageUri)
	if err != nil {
		t.Errorf("\t\tErr on GET to %v: %v", imageUri, err)
	}
	defer rsp.Body.Close()
	var actualImage []byte
	io.ReadFull(rsp.Body, actualImage)

	t.Log("\t\tThe response should be 200 OK")
	if rsp.StatusCode != 200 {
		t.Error("\t\t\tExpected 200 but was", rsp.StatusCode)
	}

	t.Log("\t\tThe Content-Type should be image/jpeg")
	if rsp.Header.Get("Content-Type") != "image/jpeg" {
		t.Error("\t\t\tExpected image/jpeg but was", rsp.Header.Get("Content-Type"))
	}

	t.Log("\t\tThe retrieved images should match each comic's first page")
	if bytes.Equal(actualImage, expectedImage) == false {
		t.Errorf("\t\t\tBytes did not match len(actual) = %v, len(expected) = %v", len(actualImage), len(expectedImage))
	}
}

func getRoot(t *testing.T, uri string) *root {
	rootRsp, err := http.Get(uri)
	if err != nil {
		t.Fatalf("Error on HTTP GET to %v: %v", uri, err)
	}
	defer rootRsp.Body.Close()
	var data root
	dec := json.NewDecoder(rootRsp.Body)
	if err := dec.Decode(&data); err != nil {
		t.Fatal(err)
	}

	return &data
}

func getFileContents(paths ...string) ([][]byte, error) {
	images := make([][]byte, 0, len(paths))
	for _, path := range paths {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		images = append(images, b)
	}

	return images, nil
}
