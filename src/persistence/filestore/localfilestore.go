package filestore

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// LocalFileStore is an implementation of the domainservices.FileStorer interface that stores
// files in a known location on disk.
type LocalFileStore struct {
	path string
}

func NewLocalFileStore(path string) *LocalFileStore {
	return &LocalFileStore{
		path: path,
	}
}

func (store *LocalFileStore) Store(key string, filenames, sourcePaths []string) error {
	keypath := store.getKeyPath(key)

	os.MkdirAll(keypath, 0700)
	for i, srcpath := range sourcePaths {
		dstpath := filepath.Join(keypath, filenames[i])
		log.Printf("copying from %v to %v\n", srcpath, dstpath)
		if err := copy(dstpath, srcpath); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (store *LocalFileStore) Open(name string) (*os.File, error) {
	p := filepath.Join(store.path, name)
	log.Printf("Opening file path %v for key %v", p, name)
	return os.Open(p)
}

func (store *LocalFileStore) GetAllKeys() (keys []string, err error) {
	d, err := os.Open(store.path)
	if err != nil {
		return nil, err
	}

	return d.Readdirnames(-1)
}

func (store *LocalFileStore) GetFilenames(key string) (names []string, err error) {
	d, err := os.Open(store.getKeyPath(key))
	if err != nil {
		return nil, err
	}

	return d.Readdirnames(-1)
}

func (store *LocalFileStore) getKeyPath(key string) string {
	return filepath.Join(store.path, key)
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
