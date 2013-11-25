package persistence

import (
	"domain/model"
)

type InMemoryEventStore struct {
	events []*model.ComicAdded // should be able to contain other types of event
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{}
}

func (store *InMemoryEventStore) AddEvent(event *model.ComicAdded) error {
	store.events = append(store.events, event)
	return nil
}

func (store *InMemoryEventStore) GetAllEvents() []*model.ComicAdded {
	return store.events
}
