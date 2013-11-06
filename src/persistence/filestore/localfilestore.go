package filestore

import (
	"io"
	"os"
	"path/filepath"
)

// localFilestore is an implementation of the domainservices.FileStorer interface that stores
// files in a known location on disk.
type localFilestore struct {
	path string
	keys []string
}

func NewLocalFileStore(path string) *localFilestore {
	return &localFilestore{
		path: path,
		keys: make([]string, 0, 50),
	}
}

func (store *localFilestore) Store(key string, files map[string]string) error {
	keypath := filepath.Join(store.path, key)

	os.MkdirAll(keypath, os.ModeDir)
	for dstfilename, srcpath := range files {
		dstpath := filepath.Join(keypath, dstfilename)
		copy(dstpath, srcpath)
	}

	return nil
}

func copy(dstpath, srcpath string) error {
	src, _ := os.Open(srcpath)
	defer src.Close()

	dst, _ := os.Create(dstpath)
	defer dst.Close()

	_, err := io.Copy(dst, src)
	return err
}
