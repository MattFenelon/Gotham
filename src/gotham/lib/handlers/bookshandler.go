package handlers

import (
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"encoding/json"
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
	metadata *addBookMetadata
	pages    map[string]string
	form     *multipart.Form
}

func (f *addBookForm) RemoveAll() error {
	if f != nil && f.form != nil {
		return f.form.RemoveAll()
	}
	return nil
}

func BooksHandler(w http.ResponseWriter, r *http.Request, eventstorer domainservices.EventStorer, filestorer domainservices.FileStorer) {
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

	comics := domainservices.NewComicDomain(eventstorer, filestorer)
	if err := comics.AddComic(uuid.NewRandom(), form.metadata.SeriesTitle, form.metadata.Title, form.pages); err != nil { // TODO: Raise different types of errors, i.e. database
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

	result.pages = getPageFiles(result.form)

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

func getPageFiles(form *multipart.Form) map[string]string {
	fileparts := form.File["page"]
	filenames := make(map[string]string, cap(fileparts))

	for _, p := range fileparts {
		filenames[p.Filename] = getFilename(p) // TODO: Error on multiple parts with the same filename
	}

	return filenames
}

func getFilename(file *multipart.FileHeader) string {
	f, _ := file.Open() // TODO: Deal with error
	defer f.Close()
	osFile := f.(*os.File)
	return osFile.Name()
}
