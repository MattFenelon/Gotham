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
