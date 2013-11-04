package handlers

import (
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"encoding/json"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

type booksPostRequest struct {
	SeriesTitle string
	Title       string
}

func BooksHandler(w http.ResponseWriter, r *http.Request, storer domainservices.EventStorer) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	// if err != nil || mediaType != "multipart/form-data" {
	// 	w.WriteHeader(http.StatusUnsupportedMediaType)
	// 	return
	// }
	boundary := params["boundary"]
	reader := multipart.NewReader(r.Body, boundary)
	// 20KB is an arbitary choice for max bytes. This API doesn't need to be quick so writing the image files
	// to disk isn't a concern and it ensures the server doesn't take up too much memory.
	form, _ := reader.ReadForm(20480)
	defer func() {
		// form.RemoveAll() // TODO: Log errors
	}()
	// if err != nil {
	// }
	// if form.Value["metadata"] == nil || len(form.Value["metadata"]) == 0 {
	// }
	// Ignore extra values or return an error?

	metadata := form.Value["metadata"][0]
	// if metadata == "" {
	// }

	metadataReader := strings.NewReader(metadata)
	dec := json.NewDecoder(metadataReader)

	var request booksPostRequest
	if err := dec.Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comics := domainservices.NewComicDomain(storer)
	if err := comics.AddComic(uuid.NewRandom(), request.SeriesTitle, request.Title); err != nil { // TODO: Raise different types of errors, i.e. database
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
