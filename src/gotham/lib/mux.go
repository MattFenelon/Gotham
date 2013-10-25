package lib

import (
	"gotham/lib/handlers"
	"io"
	"net/http"
)

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		rootHandler := handlers.NewRootHandler()
		response := rootHandler.Get()

		w.Header().Add("Content-Type", response.ContentType)
		io.WriteString(w, response.Result)
		return
	}

	http.NotFound(w, r)
}
