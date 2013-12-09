package handlers

import (
	"domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type series struct {
	Title string       `json:"title"`
	Books []seriesBook `json:"books"`
}

type seriesBook struct {
	Title         string     `json:"title"`
	PublishedDate time.Time  `json:"publishedDate"`
	WrittenBy     []string   `json:"writtenBy"`
	ArtBy         []string   `json:"artBy"`
	Blurb         string     `json:"blurb"`
	Links         []linkView `json:"links"`
}

func SeriesHandler(w http.ResponseWriter, r *http.Request, d domain.ComicDomain) {
	pathSeriesTitle := strings.TrimPrefix(r.URL.Path, "/series/")

	view := d.GetSeriesView(pathSeriesTitle)
	if view == nil {
		http.NotFound(w, r)
		return
	}

	books := make([]seriesBook, 0, len(view.Books))
	for _, b := range view.Books {
		links := []linkView{
			linkView{Rel: "self", Href: fmt.Sprintf("http://%v/books/%v", r.Host, b.Id)},
			linkView{Rel: "image", Href: fmt.Sprintf("http://%v/pages/%v", r.Host, b.ImageKey)},
		}

		book := seriesBook{
			Title:         b.Title,
			PublishedDate: b.PublishedDate,
			WrittenBy:     b.WrittenBy,
			ArtBy:         b.ArtBy,
			Blurb:         b.Blurb,
			Links:         links,
		}
		books = append(books, book)
	}
	response := series{Title: view.Title, Books: books}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(&response) // TODO: Check for errors
}
