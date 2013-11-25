package riak_tests

import (
	"domain/model"
	"persistence/riak"
	"testing"
)

func TestComicAdded(t *testing.T) {
	t.Log("When adding a single comic")

	seriesTitle, _ := model.NewSeriesTitle("Prophet")
	title, _ := model.NewBookTitle("Prophet 31")
	expectedEvent := model.NewComicAdded(model.NewRandomComicId(), seriesTitle, title, []string{"0.jpg", "1.jpg"})

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

	seriesTitle, _ := model.NewSeriesTitle("Spider-men")
	title, _ := model.NewBookTitle("Spider-men 1")
	expectedEvent1 := model.NewComicAdded(model.NewRandomComicId(), seriesTitle, title, []string{"1.jpg", "2.jpg"})
	StoreEvent(t, expectedEvent1)
	defer DeleteEvent(t, expectedEvent1.Id.String())

	seriesTitle, _ = model.NewSeriesTitle("Guardians of the Galaxy")
	title, _ = model.NewBookTitle("Guardians of the Galaxy 11")
	expectedEvent2 := model.NewComicAdded(model.NewRandomComicId(), seriesTitle, title, []string{"0.jpg"})
	StoreEvent(t, expectedEvent2)
	defer DeleteEvent(t, expectedEvent2.Id.String())

	actualEvent1 := GetEvent(t, expectedEvent1.Id.String())
	actualEvent2 := GetEvent(t, expectedEvent2.Id.String())

	t.Log("They should all be retrievable")
	if actualEvent1.Equal(expectedEvent1) == false || actualEvent2.Equal(expectedEvent2) == false {
		t.Fatalf("Expected events:\n[%v %v]\nbut actual were:\n[%v %v]", expectedEvent1, expectedEvent2, actualEvent1, actualEvent2)
	}
}

func StoreEvent(t *testing.T, event *model.ComicAdded) {
	store := riak.NewRiakEventStore(riakCluster, riakClientId)
	if err := store.AddEvent(event); err != nil {
		t.Fatal(err)
	}
}

// TODO: What type should comicId be? What's the correct interface into the model?
func GetEvent(t *testing.T, comicId string) (event *model.ComicAdded) {
	store := riak.NewRiakEventStore(riakCluster, riakClientId)
	event, err := store.GetEvent(comicId)
	if err != nil {
		t.Fatal(err)
	}
	return event
}

func DeleteEvent(t *testing.T, comicId string) {
	store := riak.NewRiakEventStore(riakCluster, riakClientId)
	if err := store.DeleteEvent(comicId); err != nil {
		t.Fatal(err)
	}
}
