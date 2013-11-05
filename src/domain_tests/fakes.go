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
	*persistence.InMemoryFileStore
}

func NewFakeFileStore() *FakeFileStore {
	return &FakeFileStore{
		InMemoryFileStore: persistence.NewInMemoryFileStore(),
	}
}
