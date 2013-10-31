package riak_tests

import (
	"domain"
	"persistence"
	"testing"
)

// TODO: Rename file?

func TestComicAdded(t *testing.T) { // TODO: Better test name
	seriesTitle, _ := domain.NewSeriesTitle("Prophet")
	title, _ := domain.NewBookTitle("Prophet 31")
	expectedEvent := domain.NewComicAdded(domain.NewRandomComicId(), seriesTitle, title)

	storeErr := StoreEvent(expectedEvent)
	if storeErr != nil {
		t.Fatalf("%v", storeErr)
	}

	actualEvent, getErr := GetEvent(expectedEvent.Id.String())
	if getErr != nil {
		t.Fatalf("%v", getErr)
	}

	if expectedEvent.Equal(actualEvent) == false {
		t.Errorf("Expected:\n%v\nbut was:\n%v", expectedEvent, actualEvent)
	}
	// TODO: Delete bucket at end of test.
}

func TestAddingEventMoreThanOnce(t *testing.T) {
	t.Skip()
}

func TestAddingMultipleEvents(t *testing.T) {
	t.Skip()
}

func StoreEvent(event *domain.ComicAdded) error {
	store := persistence.NewRiakEventStore()
	err := store.AddComic(event)
	return err
}

// TODO: What type should comicId be? What's the correct interface into the domain?
func GetEvent(comicId string) (event *domain.ComicAdded, err error) {
	store := persistence.NewRiakEventStore()
	event, err = store.GetEvent(comicId)
	return event, err
}
