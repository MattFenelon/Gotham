package lib

import (
	"net/http"
)

type Exports struct {
	Handler http.Handler
}

func Configure() (exports Exports) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", ServeHttp)

	exports = Exports{Handler: serveMux}
	return exports
}
