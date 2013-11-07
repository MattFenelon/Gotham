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

// FakeFileStore is a fake implementation of the FileStorer interface.
type FakeFileStore struct {
	stored map[string][]string
}

func NewFakeFileStore() *FakeFileStore {
	return &FakeFileStore{
		stored: map[string][]string{},
	}
}

func (store *FakeFileStore) Store(key string, filenames []string, sourcePaths []string) error {
	store.stored[key] = filenames
	return nil
}

func (store *FakeFileStore) Get(key string) []string {
	return store.stored[key]
}

// FakeViewStore
type FakeViewStore struct {
	*persistence.InMemoryViewStore
}

func NewFakeViewStore() *FakeViewStore {
	return &FakeViewStore{
		InMemoryViewStore: persistence.NewInMemoryViewStore(),
	}
}
