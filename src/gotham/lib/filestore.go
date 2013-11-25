package lib

import (
	"domain"
	"net/http"
	"os"
)

func makeFilestoreHandler(path string, filestore FileStore) http.Handler {
	return http.StripPrefix(path,
		http.FileServer(filestoreFilesystem(func(name string) (http.File, error) {
			return filestore.Open(name)
		})))
}

type FileStore interface {
	domain.FileStorer
	Open(name string) (*os.File, error)
}

type filestoreFilesystem func(name string) (http.File, error)

func (fs filestoreFilesystem) Open(name string) (http.File, error) {
	return fs(name)
}
