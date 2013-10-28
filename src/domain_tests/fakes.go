package domain_tests

import (
	"domain"
)

type FakeEventStorer struct {
	events []*domain.ComicAdded // should be able to contain other types of event
}

func NewFakeEventStorer() *FakeEventStorer {
	return &FakeEventStorer{}
}

func (repo *FakeEventStorer) AddEvent(event *domain.ComicAdded) {
	repo.events = append(repo.events, event)
}
