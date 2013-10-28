package tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"domainservices"
	"testing"
)

func TestCreateComic(t *testing.T) {
	id := uuid.NewRandom()
	eventStorer := NewFakeEventStorer()
	expected := domain.NewComicAdded(id.String(), "Prophet", "Prophet 31")

	t.Log("When adding a new comic")
	command := domainservices.NewCreateComicCommand(id, "Prophet", "Prophet 31")
	domainservices.AddComic(command, eventStorer) // TODO: Refactor to command processor

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, eventStorer)
}

func TestCreateMultipleComics(t *testing.T) {
	id := uuid.NewRandom()
	eventStorer := NewFakeEventStorer()
	expected := []*domain.ComicAdded{
		domain.NewComicAdded(id.String(), "Prophet", "Prophet 31"),
		domain.NewComicAdded(id.String(), "Batman", "Batman 1")}

	t.Log("When adding multiple comics")
	command := domainservices.NewCreateComicCommand(id, "Prophet", "Prophet 31")
	domainservices.AddComic(command, eventStorer) // TODO: Refactor to command processor

	command2 := domainservices.NewCreateComicCommand(id, "Batman", "Batman 1")
	domainservices.AddComic(command2, eventStorer) // TODO: Refactor to command processor

	t.Log("\tIt should raise a comic added event for all of the added comics")
	AssertCollectionEquality(t, expected, eventStorer)
}

type FakeEventStorer struct {
	events []*domain.ComicAdded // should be able to contain other types of event
}

func NewFakeEventStorer() *FakeEventStorer {
	return &FakeEventStorer{}
}

func (repo *FakeEventStorer) AddEvent(event *domain.ComicAdded) {
	repo.events = append(repo.events, event)
}

func AssertEquality(t *testing.T, expected *domain.ComicAdded, actual *FakeEventStorer) {
	AssertCollectionEquality(t, []*domain.ComicAdded{expected}, actual)
}

func AssertCollectionEquality(t *testing.T, expected []*domain.ComicAdded, actual *FakeEventStorer) {
	// Write tests to test the assert methods
	actualValues := make([]domain.ComicAdded, 0, len(actual.events))
	expectedValues := make([]domain.ComicAdded, 0, len(expected))

	for _, expectedEvent := range expected {
		expectedValues = append(expectedValues, *expectedEvent)
	}

	error := false
	for _, actualEvent := range actual.events {
		actual := *actualEvent
		actualValues = append(actualValues, actual)

		found := false
		for _, expectedValue := range expectedValues {
			if actual == expectedValue {
				found = true
			}
		}
		if found == false {
			error = true
		}
	}
	if error {
		t.Errorf("\tCollection was expected to only contain")
		t.Errorf("\t\t%#v\n", expectedValues)
		t.Errorf("\t\tbut contained\n")
		t.Errorf("\t\t%#v", actualValues)
	}

}
