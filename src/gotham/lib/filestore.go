package lib

import (
	"domain"
	"net/http"
	"os"
)

func makeFilestoreHandler(path string, filestore FileStore) http.Handler {
	return http.StripPrefix(path,
		http.FileServer(openFileFunc(func(name string) (http.File, error) {
			return filestore.Open(name)
		})))
}

type FileStore interface {
	domain.FileStorer
	Open(name string) (*os.File, error)
}

type openFileFunc func(name string) (http.File, error)

func (openFunc openFileFunc) Open(name string) (http.File, error) {
	return openFunc(name)
}
