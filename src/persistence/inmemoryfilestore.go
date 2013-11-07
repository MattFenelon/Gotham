package persistence

import (
	"io/ioutil"
)

type InMemoryFileStore struct {
	filenames map[string][]string
	contents  map[string][][]byte
}

func NewInMemoryFileStore() *InMemoryFileStore {
	return &InMemoryFileStore{
		filenames: map[string][]string{},
		contents:  map[string][][]byte{},
	}
}

func (store *InMemoryFileStore) Store(key string, filenames, sourcePaths []string) error {
	// TODO: Check key and files for nil
	// TODO: Check that filenames and sourcepaths have the same length

	contents := make([][]byte, 0, len(sourcePaths))
	for _, srcPath := range sourcePaths {
		bytes, err := ioutil.ReadFile(srcPath)
		if err != nil {
			return err
		}
		contents = append(contents, bytes)
	}

	store.filenames[key] = filenames
	store.contents[key] = contents

	return nil
}

func (f *InMemoryFileStore) Get(key string) []string {
	return f.filenames[key]
}

func (f *InMemoryFileStore) GetAll() (files map[string][]string, contents map[string][][]byte) {
	return f.filenames, f.contents
}
