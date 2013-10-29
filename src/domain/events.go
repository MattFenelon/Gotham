package domain

import (
	"fmt"
)

type ComicAdded struct {
	Id          comicId
	SeriesTitle seriesTitle
	Title   bookTitle
}

func NewComicAdded(comicId comicId, seriesTitle seriesTitle, bookTitle bookTitle) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, Title: bookTitle}
}

func (a ComicAdded) Equal(b interface{}) bool {
	if v, ok := b.(ComicAdded); ok {
		return a.EqualTo(v)
	}

	if p, ok := b.(*ComicAdded); ok {
		return a.EqualTo(*p)
	}

	return false
}

func (a ComicAdded) EqualTo(b ComicAdded) bool {
	return a.Id.Equal(b.Id) && a.SeriesTitle == b.SeriesTitle && a.Title == b.Title
}

func (ca ComicAdded) String() string {
	return fmt.Sprintf("%T:{id: %v, seriesTitle: \"%v\", bookTitle: \"%v\"}", ca, ca.Id, ca.SeriesTitle, ca.Title)
}
