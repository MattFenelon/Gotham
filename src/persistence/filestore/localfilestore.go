package filestore

import (
	"io"
	"log"
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

func (store *localFilestore) Store(key string, filenames, sourcePaths []string) error {
	keypath := filepath.Join(store.path, key)

	os.MkdirAll(keypath, os.ModeDir)
	for i, srcpath := range sourcePaths {
		dstpath := filepath.Join(keypath, filenames[i])
		log.Printf("copying from %v to %v\n", srcpath, dstpath)
		if err := copy(dstpath, srcpath); err != nil {
			return err
		}
	}

	return nil
}

func copy(dstpath, srcpath string) error {
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
