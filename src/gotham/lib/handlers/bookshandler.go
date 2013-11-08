package handlers

import (
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"encoding/json"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type addBookMetadata struct {
	SeriesTitle string
	Title       string
}

type addBookForm struct {
	metadata      *addBookMetadata
	pageFilenames []string
	pageSources   []string
	form          *multipart.Form
}

func (f *addBookForm) RemoveAll() error {
	if f != nil && f.form != nil {
		return f.form.RemoveAll()
	}
	return nil
}

func BooksHandler(w http.ResponseWriter, r *http.Request, d domainservices.ComicDomain) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	form, errstatuscode := readForm(r)
	defer form.RemoveAll() // TODO: Deal with errors
	if errstatuscode != 0 {
		w.WriteHeader(errstatuscode)
		return
	}

	if err := d.AddComic(uuid.NewRandom(), form.metadata.SeriesTitle, form.metadata.Title, form.pageFilenames, form.pageSources); err != nil { // TODO: Raise different types of errors, i.e. database
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func readForm(r *http.Request) (result *addBookForm, errstatuscode int) {
	result = &addBookForm{}

	boundary, errstatuscode := getBoundary(r)
	if errstatuscode != 0 {
		return nil, errstatuscode
	}

	reader := multipart.NewReader(r.Body, boundary)
	// -1 forces the files to be written to disk.
	// This API doesn't need to be quick so writing the image files
	// to disk isn't a concern and it ensures the server doesn't take up too much memory.
	// It also means the form.Files can be cast to *os.File.
	var err error
	result.form, err = reader.ReadForm(-1)
	if err != nil {
		return nil, http.StatusBadRequest
	}

	result.metadata, errstatuscode = getMetadata(result.form)
	if errstatuscode != 0 {
		return result, errstatuscode
	}

	result.pageFilenames, result.pageSources = getPageFiles(result.form)

	return result, 0
}

func getBoundary(r *http.Request) (boundary string, statuscodeerr int) {
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "multipart/form-data" {
		return "", http.StatusUnsupportedMediaType
	}

	boundary = params["boundary"]
	if boundary == "" {
		return "", http.StatusUnsupportedMediaType
	}

	return boundary, 0
}

func getMetadata(form *multipart.Form) (metadata *addBookMetadata, errstatuscode int) {
	values := form.Value["metadata"]
	if len(values) == 0 {
		return nil, http.StatusBadRequest
	}

	dec := json.NewDecoder(strings.NewReader(values[0]))
	if err := dec.Decode(&metadata); err != nil {
		return nil, http.StatusBadRequest
	}

	return metadata, 0
}

func getPageFiles(form *multipart.Form) (filenames, sources []string) {
	fileparts := form.File["page"]
	log.Printf("Received form contains %v page images", len(fileparts))
	filenames = make([]string, 0, len(fileparts))
	sources = make([]string, 0, len(fileparts))

	for _, p := range fileparts {
		log.Printf("Received page image %v", p.Filename)
		filenames = append(filenames, p.Filename) // TODO: Error on multiple parts with the same filename
		sources = append(sources, getFilename(p))
	}

	return filenames, sources
}

func getFilename(file *multipart.FileHeader) string {
	f, _ := file.Open() // TODO: Deal with error
	defer f.Close()
	osFile := f.(*os.File)
	return osFile.Name()
}
