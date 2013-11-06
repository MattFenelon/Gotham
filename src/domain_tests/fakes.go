package domain_tests

import (
	"persistence"
)

type FakeEventStorer struct {
	*persistence.InMemoryEventStore
}

func NewFakeEventStorer() *FakeEventStorer {
	return &FakeEventStorer{InMemoryEventStore: persistence.NewInMemoryEventStore()}
}

type FakeFileStore struct {
	stored map[string][]string
}

func NewFakeFileStore() *FakeFileStore {
	return &FakeFileStore{
		stored: map[string][]string{},
	}
}

func (store *FakeFileStore) Store(key string, files map[string]string) error {
	names := make([]string, 0, len(files))
	for f, _ := range files {
		names = append(names, f)
	}
	store.stored[key] = names
	return nil
}

func (store *FakeFileStore) Get(key string) []string {
	return store.stored[key]
}

func (store *FakeFileStore) GetAll() map[string][]string {
	return store.stored
}
