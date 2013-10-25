package tests

import (
	"code.google.com/p/go-uuid/uuid"
	"testing"
)

func TestCreateComic(t *testing.T) {
	id := uuid.NewRandom()
	comicEvents := []*ComicAdded{} // TODO: repo
	expected := NewComicAdded(id.String(), "Prophet", "Prophet 31")

	t.Log("When adding a new comic")
	command := &CreateComicCommand{comicId: id, seriesTitle: "Prophet", bookTitle: "Prophet 31"}
	AddComic(command, &comicEvents)

	t.Log("\tIt should raise a comic added event")
	AssertEquality(t, expected, comicEvents)
}

type CreateComicCommand struct {
	comicId     uuid.UUID
	seriesTitle string
	bookTitle   string
}

func AddComic(command *CreateComicCommand, events *[]*ComicAdded) { // events should contain other events
	event := NewComicAdded(command.comicId.String(), command.seriesTitle, command.bookTitle)
	addedEvents := append(*events, event)
	events = &addedEvents
}

func NewComicAdded(comicId string, seriesTitle, bookTitle string) *ComicAdded {
	return &ComicAdded{id: comicId, seriesTitle: seriesTitle, bookTitle: bookTitle}
}

type ComicAdded struct {
	id          string // TODO: Create an identifier type
	seriesTitle string
	bookTitle   string
}

func AssertEquality(t *testing.T, expected *ComicAdded, actual []*ComicAdded) {
	values := make([]ComicAdded, 0, len(actual))
	error := false
	for _, event := range actual {
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
