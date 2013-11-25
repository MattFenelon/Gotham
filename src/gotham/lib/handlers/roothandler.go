package handlers

import (
	"domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type rootView struct {
	Series []rootViewSeries `json:"series"`
}

type rootViewSeries struct {
	Title string     `json:"title"`
	Links []linkView `json:"links"`
}

func RootHandler(w http.ResponseWriter, r *http.Request, d domain.ComicDomain) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	src := d.GetFrontPageView()
	dst := rootView{
		Series: make([]rootViewSeries, 0, len(src.Series)),
	}

	for _, s := range src.Series {
		series := rootViewSeries{
			Title: s.Title,
			Links: []linkView{
				linkView{Rel: "seriesimage", Href: fmt.Sprintf("http://%v/pages/%v", r.Host, s.ImageKey)},
				linkView{Rel: "promotedbook", Href: fmt.Sprintf("http://%v/books/%v", r.Host, s.PromotedBookId)},
			},
		}
		dst.Series = append(dst.Series, series)
	}

	enc := json.NewEncoder(w)
	enc.Encode(dst)
}
