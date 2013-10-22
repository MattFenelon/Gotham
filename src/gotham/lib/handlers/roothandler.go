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
	"set": [
		{
			"type": "series",
			"title": "Prophet",
			"links": {
				"via": {"href": "/series/1"},
				"{docHost}/rel/seriesimage": {"href": "/images/1.jpg"}
			},
		},
		{
			"type": "series",
			"title": "Jupiter's Legacy",
			"links": {
				"via": {"href": "/series/2"},
				"{docHost}/rel/seriesimage": {"href": "/images/2.jpg"}
			},
		}
	]
}`}

	return result
}
