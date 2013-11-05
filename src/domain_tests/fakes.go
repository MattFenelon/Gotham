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
		stored: make(map[string][]string),
	}
}

func (f *FakeFileStore) Store(key string, filepaths []string) error {
	f.stored[key] = filepaths
	return nil
}

func (f *FakeFileStore) GetAll(key string) []string {
	return f.stored[key]
}
