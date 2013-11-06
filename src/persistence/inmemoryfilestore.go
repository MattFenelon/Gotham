package persistence

import (
	"io/ioutil"
)

type InMemoryFileStore struct {
	stored map[string]map[string][]byte
}

func NewInMemoryFileStore() *InMemoryFileStore {
	return &InMemoryFileStore{
		stored: map[string]map[string][]byte{},
	}
}

func (store *InMemoryFileStore) Store(key string, files map[string]string) error {
	// TODO: Check key and files for nil

	f := make(map[string][]byte, len(files))
	for filename, path := range files {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		f[filename] = contents
	}
	store.stored[key] = f
	return nil
}

func (f *InMemoryFileStore) Get(key string) []string {
	files := f.stored[key]
	names := make([]string, 0, len(files))
	for filename, _ := range files {
		names = append(names, filename)
	}
	return names
}

func (f *InMemoryFileStore) GetAll() map[string]map[string][]byte {
	return f.stored
}
