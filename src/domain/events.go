package domain

import (
	"fmt"
)

type ComicAdded struct {
	Id          comicId // TODO: Create an identifier type
	SeriesTitle seriesTitle
	BookTitle   bookTitle
}

func NewComicAdded(comicId comicId, seriesTitle seriesTitle, bookTitle bookTitle) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, BookTitle: bookTitle}
}

func (a ComicAdded) Equal(b interface{}) bool {
	if r, ok := b.(ComicAdded); ok {
		return a.EqualTo(r)
	}

	return false
}

func (a ComicAdded) EqualTo(b ComicAdded) bool {
	return a.Id.Equal(b.Id) && a.SeriesTitle == b.SeriesTitle && a.BookTitle == b.BookTitle
}

func (ca ComicAdded) String() string {
	return fmt.Sprintf("%T:{id: %v, seriesTitle: \"%v\", bookTitle: \"%v\"}", ca, ca.Id, ca.SeriesTitle, ca.BookTitle)
}
