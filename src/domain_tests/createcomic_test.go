package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"domainservices"
	"testing"
)

func TestCreateComic(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	id := uuid.NewRandom()
	expected := domain.NewComicAdded(id.String(), "Prophet", "Prophet 31")

	t.Log("When adding a new comic")
	comics.AddComic(id, "Prophet", "Prophet 31")

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, eventStorer)
}

func TestCreateMultipleComics(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()
	expected := []*domain.ComicAdded{
		domain.NewComicAdded(id1.String(), "Prophet", "Prophet 31"),
		domain.NewComicAdded(id2.String(), "Batman", "Batman 1")}

	t.Log("When adding multiple comics")
	comics.AddComic(id1, "Prophet", "Prophet 31")
	comics.AddComic(id2, "Batman", "Batman 1")

	t.Log("\tIt should raise a comic added event for all of the added comics")
	AssertCollectionEquality(t, expected, eventStorer)
}

// TODO: Test adding multiple comics with same id.
