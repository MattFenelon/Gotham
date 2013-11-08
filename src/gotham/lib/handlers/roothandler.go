package handlers

import (
	"io"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	result := `{
	"series": [
		{
			"title": "Prophet",
			"links": {
				"seriesimage": {"href": "http://gotham/images/0.jpg"}
			},
		},
		{
			"title": "Jupiter's Legacy",
			"links": {
				"seriesimage": {"href": "http://gotham/images/0.jpg"}
			},
		}
	]
}`

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, result)
}
