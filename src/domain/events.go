package domain

import (
	"fmt"
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
	if a.Id.Equal(b.Id) && a.SeriesTitle == b.SeriesTitle && a.Title == b.Title {
		if len(a.Pages) != len(b.Pages) {
			return false
		}
		for i, v := range a.Pages {
			if b.Pages[i] != v {
				return false
			}
		}
		return true
	}

	return false
}

func (ca ComicAdded) String() string {
	return fmt.Sprintf("%T:{id: %v, seriesTitle: \"%v\", bookTitle: \"%v\", pages: %v}", ca, ca.Id, ca.SeriesTitle, ca.Title, ca.Pages)
}
