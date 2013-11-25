package handlers

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type comic struct {
	Links []linkView `json:"links"`
}

func BookHandler(w http.ResponseWriter, r *http.Request, d domain.ComicDomain) {
	pathId := strings.TrimPrefix(r.URL.Path, "/books/")
	if pathId == r.URL.Path {
		http.NotFound(w, r)
		return
	}

	id := uuid.Parse(pathId)
	view := d.GetComicView(id)
	if view == nil {
		http.NotFound(w, r)
		return
	}

	links := make([]linkView, 0, len(view.Pages))
	for _, p := range view.Pages {
		links = append(links, linkView{Rel: "item", Href: fmt.Sprintf("http://%v/pages/%v/%v", r.Host, id.String(), p)})
	}

	response := comic{Links: links}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(&response) // TODO: Check for errors
}
