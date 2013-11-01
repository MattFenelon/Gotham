package handlers

import (
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"encoding/json"
	"net/http"
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

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

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
