package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"domainservices"
	"testing"
)

func TestCreateComic(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	fileStorer := NewFakeFileStore()
	comics := domainservices.NewComicDomain(eventStorer, fileStorer)

	id := uuid.NewRandom()
	expectedPages := []string{"0.jpg", "1.jpg", "2.jpg", "3.jpg"}
	expected := NewComicAdded(id, "Prophet", "Prophet 31", expectedPages)

	t.Log("When adding a new comic")
	comics.AddComic(id, "Prophet", "Prophet 31", []string{"0.jpg", "1.jpg", "2.jpg", "3.jpg"})

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, eventStorer.GetAllEvents())

	t.Log("\tIt should persist the page images")
	AssertCollectionEquality(t, expectedPages, fileStorer.Get(id.String()))
}

func TestCreateMultipleComics(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	fileStorer := NewFakeFileStore()
	comics := domainservices.NewComicDomain(eventStorer, fileStorer)

	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()
	expectedPages1 := []string{"0.jpg", "1.jpg", "2.jpg"}
	expectedPages2 := []string{"0.jpg"}
	expected := []*domain.ComicAdded{
		NewComicAdded(id1, "Prophet", "Prophet 31", expectedPages1),
		NewComicAdded(id2, "Batman", "Batman 1", expectedPages2)}

	t.Log("When adding multiple comics")
	comics.AddComic(id1, "Prophet", "Prophet 31", []string{"0.jpg", "1.jpg", "2.jpg"})
	comics.AddComic(id2, "Batman", "Batman 1", []string{"0.jpg"})

	t.Log("\tIt should raise a comic added event for all of the added comics")
	AssertCollectionEquality(t, expected, eventStorer.GetAllEvents())

	t.Log("\tIt should persist the page images for each comic")
	AssertCollectionEquality(t, expectedPages1, fileStorer.Get(id1.String()))
	AssertCollectionEquality(t, expectedPages2, fileStorer.Get(id2.String()))
}

func TestCreateComicTitleTrimming(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	fileStorer := NewFakeFileStore()
	comics := domainservices.NewComicDomain(eventStorer, fileStorer)

	id := uuid.NewRandom()
	expected := NewComicAdded(id, "Series With Whitespace", "Title With Whitespace", []string{})

	t.Log("When adding a new comic with whitespace in the book title and series title")
	comics.AddComic(id, "\t\n\v\f\r\u0085\u00A0Series With Whitespace\t\n\v\f\r\u0085\u00A0", "\t\n\v\f\r\u0085\u00A0Title With Whitespace\t\n\v\f\r\u0085\u00A0", []string{})

	t.Log("\tIt should remove the extra whitespace")
	AssertEquality(t, expected, eventStorer.GetAllEvents())
}

func TestCreateComicNoBookTitle(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	fileStorer := NewFakeFileStore()
	comics := domainservices.NewComicDomain(eventStorer, fileStorer)

	t.Log("When the book title is an empty string\nor is only whitespace")
	emptyErr := comics.AddComic(uuid.NewRandom(), "Batman & Robin", "", []string{})
	whitespaceErr := comics.AddComic(uuid.NewRandom(), "Batman & Robin", "\t\n\v\f\r\u0085\u00A0", []string{})

	t.Log("\tFor the empty string it should return an error specifying that a book title is required")
	if emptyErr == nil || emptyErr.Error() != "Book title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tFor the whitespace string it should return an error specifying that a book title is required")
	if whitespaceErr == nil || whitespaceErr.Error() != "Book title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tIt should not add the comics")
	AssertCollectionIsEmpty(t, eventStorer.GetAllEvents())
}

func TestCreateComicNoSeriesTitle(t *testing.T) {
	eventStorer := NewFakeEventStorer()
	fileStorer := NewFakeFileStore()
	comics := domainservices.NewComicDomain(eventStorer, fileStorer)

	t.Log("When the series title is an empty string\nor is only whitespace")
	emptyErr := comics.AddComic(uuid.NewRandom(), "", "Batman 99", []string{})
	whitespaceErr := comics.AddComic(uuid.NewRandom(), "\t\n\v\f\r\u0085\u00A0", "Batman 99", []string{})

	t.Log("\tFor the empty string it should return an error specifying that a series title is required")
	if emptyErr == nil || emptyErr.Error() != "Series title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tFor the whitespace string it should return an error specifying that a series title is required")
	if whitespaceErr == nil || whitespaceErr.Error() != "Series title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tIt should not add the comics")
	AssertCollectionIsEmpty(t, eventStorer.GetAllEvents())
}

// TODO: Test adding multiple comics with same id.
