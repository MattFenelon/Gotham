package model

import (
	"fmt"
	"reflect"
)

type ComicAdded struct {
	Id          comicId
	SeriesTitle seriesTitle
	Title       bookTitle
	Pages       []string
}

func NewComicAdded(comicId comicId, seriesTitle seriesTitle, bookTitle bookTitle, pages []string) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, Title: bookTitle, Pages: pages}
}

func (a ComicAdded) Equal(b interface{}) bool {
	if v, ok := b.(ComicAdded); ok {
		return a.equalTo(&v)
	}

	if p, ok := b.(*ComicAdded); ok {
		return a.equalTo(p)
	}

	return false
}

func (a *ComicAdded) equalTo(b *ComicAdded) bool {
	return reflect.DeepEqual(a, b)
}

func (ca ComicAdded) String() string {
	return fmt.Sprintf("%T:{id: %v, seriesTitle: \"%v\", bookTitle: \"%v\", pages: %v}", ca, ca.Id, ca.SeriesTitle, ca.Title, ca.Pages)
}
