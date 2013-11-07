package handlers

import (
	"gotham/lib/http"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (handler *RootHandler) Get() (result *http.HttpResponse) {
	result = &http.HttpResponse{
		ContentType: "application/json",
		Result: `{
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
}`}

	return result
}
