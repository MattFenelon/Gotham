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
	expected := domain.NewComicAdded(id.String(), domain.NewTrimmedString("Prophet"), domain.NewTrimmedString("Prophet 31"))

	t.Log("When adding a new comic")
	comics.AddComic(id, "Prophet", "Prophet 31")

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, eventStorer)
}

func TestCreateComicTitleTrimming(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	id := uuid.NewRandom()
	expected := domain.NewComicAdded(id.String(), domain.NewTrimmedString("Series With Whitespace"), domain.NewTrimmedString("Title With Whitespace"))

	t.Log("When adding a new comic with whitespace in the book title and series title")
	comics.AddComic(id, "\t\n\v\f\r\u0085\u00A0Series With Whitespace\t\n\v\f\r\u0085\u00A0", "\t\n\v\f\r\u0085\u00A0Title With Whitespace\t\n\v\f\r\u0085\u00A0")

	t.Log("\tIt should remove the extra whitespace")
	AssertEquality(t, expected, eventStorer)
}

func TestCreateMultipleComics(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	comics := domainservices.NewComicDomain(eventStorer)

	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()
	expected := []*domain.ComicAdded{
		domain.NewComicAdded(id1.String(), domain.NewTrimmedString("Prophet"), domain.NewTrimmedString("Prophet 31")),
		domain.NewComicAdded(id2.String(), domain.NewTrimmedString("Batman"), domain.NewTrimmedString("Batman 1"))}

	t.Log("When adding multiple comics")
	comics.AddComic(id1, "Prophet", "Prophet 31")
	comics.AddComic(id2, "Batman", "Batman 1")

	t.Log("\tIt should raise a comic added event for all of the added comics")
	AssertCollectionEquality(t, expected, eventStorer)
}

// TODO: Test adding multiple comics with same id.
