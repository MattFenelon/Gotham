package persistence

import (
	"domain"
)

type InMemoryEventStore struct {
	events []*domain.ComicAdded // should be able to contain other types of event
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{}
}

func (store *InMemoryEventStore) AddEvent(event *domain.ComicAdded) error {
	store.events = append(store.events, event)
	return nil
}

func (store *InMemoryEventStore) GetAllEvents() []*domain.ComicAdded {
	return store.events
}
