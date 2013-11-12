package lib

import (
	"domainservices"
	"net/http"
	"os"
)

func makeFilestoreHandler(path string, filestore *FileStore) {
	return http.StripPrefix(path,
		http.FileServer(filestoreFilesystem(func(name string) (http.File, error) {
			return filestore.Open(name)
		})))
}

type FileStore interface {
	domainservices.FileStorer
	Open(name string) (*os.File, error)
}

type filestoreFilesystem func(name string) (http.File, error)

func (fs filestoreFilesystem) Open(name string) (http.File, error) {
	return fs(name)
}
