package tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"domainservices"
	"testing"
)

func TestCreateComic(t *testing.T) {
	id := uuid.NewRandom()
	repo := NewFakeComicRepository()
	expected := domain.NewComicAdded(id.String(), "Prophet", "Prophet 31")

	t.Log("When adding a new comic")
	command := domainservices.NewCreateComicCommand(id, "Prophet", "Prophet 31")
	domainservices.AddComic(command, repo) // TODO: Refactor to command processor

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, repo)
}

type FakeComicRepository struct {
	events []*domain.ComicAdded // should be able to contain other types of event
}

func NewFakeComicRepository() *FakeComicRepository {
	return &FakeComicRepository{}
}

func (repo *FakeComicRepository) AddEvent(event *domain.ComicAdded) {
	repo.events = append(repo.events, event)
}

func AssertEquality(t *testing.T, expected *domain.ComicAdded, actual *FakeComicRepository) {
	values := make([]domain.ComicAdded, 0, len(actual.events))
	error := false
	for _, event := range actual.events {
		values = append(values, *event)
		if *event != *expected {
			error = true
		}
	}
	if error {
		t.Errorf("\tCollection was expected to only contain")
		t.Errorf("\t\t%#v\n", expected)
		t.Errorf("\t\tbut contained\n")
		t.Errorf("\t\t%#v", values)
	}
}
