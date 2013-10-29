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
	expected := NewComicAdded(id, "Prophet", "Prophet 31")

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
		NewComicAdded(id1, "Prophet", "Prophet 31"),
		NewComicAdded(id2, "Batman", "Batman 1")}

	t.Log("When adding multiple comics")
	comics.AddComic(id1, "Prophet", "Prophet 31")
	comics.AddComic(id2, "Batman", "Batman 1")

	t.Log("\tIt should raise a comic added event for all of the added comics")
	AssertCollectionEquality(t, expected, eventStorer)
}

func TestCreateComicTitleTrimming(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	id := uuid.NewRandom()
	expected := NewComicAdded(id, "Series With Whitespace", "Title With Whitespace")

	t.Log("When adding a new comic with whitespace in the book title and series title")
	comics.AddComic(id, "\t\n\v\f\r\u0085\u00A0Series With Whitespace\t\n\v\f\r\u0085\u00A0", "\t\n\v\f\r\u0085\u00A0Title With Whitespace\t\n\v\f\r\u0085\u00A0")

	t.Log("\tIt should remove the extra whitespace")
	AssertEquality(t, expected, eventStorer)
}

func TestCreateComicNoBookTitle(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	t.Log("When adding a new comic without a book title")
	err := comics.AddComic(uuid.NewRandom(), "Batman & Robin", "")

	t.Log("\tIt should return an error specifying that a book title is required")
	if err == nil || err.Error() != "Book title's cannot be empty" {
		t.Errorf("\t\tError was %#v", err)
	}

	t.Log("\tIt should not add the comic")
	AssertCollectionIsEmpty(t, eventStorer)
}

func TestCreateComicNoSeriesTitle(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	t.Log("When adding a new comic without a book title")
	err := comics.AddComic(uuid.NewRandom(), "", "Batman 99")

	t.Log("\tIt should return an error specifying that a series title is required")
	if err == nil || err.Error() != "Series title's cannot be empty" {
		t.Errorf("\t\tError was %#v", err)
	}

	t.Log("\tIt should not add the comic")
	AssertCollectionIsEmpty(t, eventStorer)
}

// TODO: Test adding multiple comics with same id.
