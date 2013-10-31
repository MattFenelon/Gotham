package riak_tests

import (
	"domain"
	"persistence"
	"testing"
)

// TODO: Rename file?

func TestComicAdded(t *testing.T) { // TODO: Better test name
	t.Log("When adding a single comic")

	seriesTitle, _ := domain.NewSeriesTitle("Prophet")
	title, _ := domain.NewBookTitle("Prophet 31")
	expectedEvent := domain.NewComicAdded(domain.NewRandomComicId(), seriesTitle, title)

	StoreEvent(t, expectedEvent)
	defer DeleteEvent(t, expectedEvent.Id.String())

	actualEvent := GetEvent(t, expectedEvent.Id.String())

	t.Log("It should be retrievable")
	if expectedEvent.Equal(actualEvent) == false {
		t.Errorf("Expected:\n%v\nbut was:\n%v", expectedEvent, actualEvent)
	}
}

func TestAddingMultipleEvents(t *testing.T) {
	t.Log("When adding multiple comics")

	seriesTitle, _ := domain.NewSeriesTitle("Spider-men")
	title, _ := domain.NewBookTitle("Spider-men 1")
	expectedEvent1 := domain.NewComicAdded(domain.NewRandomComicId(), seriesTitle, title)
	StoreEvent(t, expectedEvent1)
	defer DeleteEvent(t, expectedEvent1.Id.String())

	seriesTitle, _ = domain.NewSeriesTitle("Guardians of the Galaxy")
	title, _ = domain.NewBookTitle("Guardians of the Galaxy 11")
	expectedEvent2 := domain.NewComicAdded(domain.NewRandomComicId(), seriesTitle, title)
	StoreEvent(t, expectedEvent2)
	defer DeleteEvent(t, expectedEvent2.Id.String())

	actualEvent1 := GetEvent(t, expectedEvent1.Id.String())
	actualEvent2 := GetEvent(t, expectedEvent2.Id.String())

	t.Log("They should all be retrievable")
	if actualEvent1.Equal(expectedEvent1) == false || actualEvent2.Equal(expectedEvent2) == false {
		t.Fatalf("Expected events:\n[%v %v]\nbut actual were:\n[%v %v]", expectedEvent1, expectedEvent2, actualEvent1, actualEvent2)
	}
}

func StoreEvent(t *testing.T, event *domain.ComicAdded) {
	store := persistence.NewRiakEventStore()
	if err := store.AddComic(event); err != nil {
		t.Fatal(err)
	}
}

// TODO: What type should comicId be? What's the correct interface into the domain?
func GetEvent(t *testing.T, comicId string) (event *domain.ComicAdded) {
	store := persistence.NewRiakEventStore()
	event, err := store.GetEvent(comicId)
	if err != nil {
		t.Fatal(err)
	}
	return event
}

func DeleteEvent(t *testing.T, comicId string) {
	store := persistence.NewRiakEventStore()
	if err := store.DeleteEvent(comicId); err != nil {
		t.Fatal(err)
	}
}
