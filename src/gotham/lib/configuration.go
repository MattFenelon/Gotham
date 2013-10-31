package lib

import (
	"domainservices"
	"net/http"
)

type Exports struct {
	Handler http.Handler
}

func Configure(store domainservices.EventStorer) (exports Exports) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", ServeHttp)

	exports = Exports{Handler: serveMux}
	return exports
}
