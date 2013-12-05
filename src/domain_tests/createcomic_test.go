package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"domain/model"
	"testing"
	"time"
)

func TestCreateComic(t *testing.T) {
	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	comics := domain.NewComicDomain(es, fs, vs)

	id := uuid.NewRandom()
	expectedPages := []string{"0.jpg", "1.jpg", "2.jpg", "3.jpg"}
	expected := NewComicAdded(
		id,
		"Prophet",
		"Prophet 31",
		[]string{"Brandon Graham", "Simon Roy", "Giannis Milonogiannis"},
		[]string{"Giannis Milonogiannis"},
		expectedPages,
		time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC),
		"Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance.")

	t.Log("When adding a new comic")
	comics.AddComic(
		id,
		"Prophet",
		"Prophet 31",
		[]string{"0.jpg", "1.jpg", "2.jpg", "3.jpg"},
		[]string{"\\source\\path\\to\\0.jpg", "\\source\\path\\to\\1.jpg", "\\source\\path\\to\\2.jpg", "\\source\\path\\to\\3.jpg"},
		[]string{"Brandon Graham", "Simon Roy", "Giannis Milonogiannis"},
		[]string{"Giannis Milonogiannis"},
		time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC),
		"Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance.")

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, es.GetAllEvents())

	t.Log("\tIt should persist the page images")
	AssertCollectionEquality(t, expectedPages, fs.Get(id.String()))
}

func TestCreateMultipleComics(t *testing.T) {
	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	comics := domain.NewComicDomain(es, fs, vs)

	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()
	expectedPages1 := []string{"0.jpg", "1.jpg", "2.jpg"}
	expectedPages2 := []string{"0.jpg"}
	expected := []*model.ComicAdded{
		NewComicAdded(id1, "Prophet", "Prophet 31", []string{"Brandon Graham"}, []string{"Giannis Milonogiannis"}, expectedPages1, time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC), "Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance."),
		NewComicAdded(id2, "Batman", "Batman 1", []string{"Bob Kane"}, []string{}, expectedPages2, time.Date(1940, time.April, 1, 0, 0, 0, 0, time.UTC), "The Legend of the Batman - Who He is, and How he Came to Be"),
	}

	t.Log("When adding multiple comics")
	comics.AddComic(id1, "Prophet", "Prophet 31", []string{"0.jpg", "1.jpg", "2.jpg"}, []string{"\\source\\path\\to\\0.jpg", "\\source\\path\\to\\1.jpg", "\\source\\path\\to\\2.jpg"}, []string{"Brandon Graham"}, []string{"Giannis Milonogiannis"}, time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC), "Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance.")
	comics.AddComic(id2, "Batman", "Batman 1", []string{"0.jpg"}, []string{"\\source\\path\\to\\0.jpg"}, []string{"Bob Kane"}, []string{}, time.Date(1940, time.April, 1, 0, 0, 0, 0, time.UTC), "The Legend of the Batman - Who He is, and How he Came to Be")

	t.Log("\tIt should raise a comic added event for all of the added comics")
	AssertCollectionEquality(t, expected, es.GetAllEvents())

	t.Log("\tIt should persist the page images for each comic")
	AssertCollectionEquality(t, expectedPages1, fs.Get(id1.String()))
	AssertCollectionEquality(t, expectedPages2, fs.Get(id2.String()))
}

func TestCreateComicTitleTrimming(t *testing.T) {
	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	comics := domain.NewComicDomain(es, fs, vs)

	id := uuid.NewRandom()
	expected := NewComicAdded(id, "Series With Whitespace", "Title With Whitespace", []string{}, []string{}, []string{"0.jpg"}, time.Time{}, "")

	t.Log("When adding a new comic with whitespace in the book title and series title")
	comics.AddComic(id, "\t\n\v\f\r\u0085\u00A0Series With Whitespace\t\n\v\f\r\u0085\u00A0", "\t\n\v\f\r\u0085\u00A0Title With Whitespace\t\n\v\f\r\u0085\u00A0", []string{"0.jpg"}, []string{"//source//0.jpg"}, []string{}, []string{}, time.Time{}, "")

	t.Log("\tIt should remove the extra whitespace")
	AssertEquality(t, expected, es.GetAllEvents())
}

func TestCreateComicNoBookTitle(t *testing.T) {
	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	comics := domain.NewComicDomain(es, fs, vs)

	t.Log("When the book title is an empty string\nor is only whitespace")
	emptyErr := comics.AddComic(uuid.NewRandom(), "Batman & Robin", "", []string{}, []string{}, []string{}, []string{}, time.Time{}, "")
	whitespaceErr := comics.AddComic(uuid.NewRandom(), "Batman & Robin", "\t\n\v\f\r\u0085\u00A0", []string{}, []string{}, []string{}, []string{}, time.Time{}, "")

	t.Log("\tFor the empty string it should return an error specifying that a book title is required")
	if emptyErr == nil || emptyErr.Error() != "Book title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tFor the whitespace string it should return an error specifying that a book title is required")
	if whitespaceErr == nil || whitespaceErr.Error() != "Book title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tIt should not add the comics")
	AssertCollectionIsEmpty(t, es.GetAllEvents())
}

func TestCreateComicNoSeriesTitle(t *testing.T) {
	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	comics := domain.NewComicDomain(es, fs, vs)

	t.Log("When the series title is an empty string\nor is only whitespace")
	emptyErr := comics.AddComic(uuid.NewRandom(), "", "Batman 99", []string{}, []string{}, []string{}, []string{}, time.Time{}, "")
	whitespaceErr := comics.AddComic(uuid.NewRandom(), "\t\n\v\f\r\u0085\u00A0", "Batman 99", []string{}, []string{}, []string{}, []string{}, time.Time{}, "")

	t.Log("\tFor the empty string it should return an error specifying that a series title is required")
	if emptyErr == nil || emptyErr.Error() != "Series title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tFor the whitespace string it should return an error specifying that a series title is required")
	if whitespaceErr == nil || whitespaceErr.Error() != "Series title cannot be empty" {
		t.Errorf("\t\tError was %#v", emptyErr)
	}

	t.Log("\tIt should not add the comics")
	AssertCollectionIsEmpty(t, es.GetAllEvents())
}

func TestCreateComicNoPages(t *testing.T) {
	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	comics := domain.NewComicDomain(es, fs, vs)

	t.Log("When the comic has no pages, or page sources")
	pagesErr := comics.AddComic(uuid.NewRandom(), "Batman", "Batman 99", []string{}, []string{}, []string{}, []string{}, time.Time{}, "")
	sourcesErr := comics.AddComic(uuid.NewRandom(), "Batman", "Batman 99", []string{"0.jpg"}, []string{}, []string{}, []string{}, time.Time{}, "")

	t.Log("\tIt should return an error specifying that pages are required")
	if pagesErr == nil || pagesErr.Error() != "At least one page is required" {
		t.Errorf("\t\tError was %#v", pagesErr)
	}

	t.Log("\tIt should return an error specifying that page sources are required")
	if sourcesErr == nil || sourcesErr.Error() != "At least one page source is required" {
		t.Errorf("\t\tError was %#v", pagesErr)
	}

	t.Log("\tIt should not add the comics")
	AssertCollectionIsEmpty(t, es.GetAllEvents())
}

// TODO: Test adding multiple comics with same id.
