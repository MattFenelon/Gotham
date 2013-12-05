package model

import (
	"fmt"
	"reflect"
	"time"
)

type ComicAdded struct {
	Id            comicId
	SeriesTitle   seriesTitle
	Title         bookTitle
	PublishedDate time.Time
	WrittenBy     []string
	ArtBy         []string
	Blurb         string
	Pages         []string
}

func NewComicAdded(comicId comicId, seriesTitle seriesTitle, bookTitle bookTitle, writtenBy, artBy, pages []string, publishedDate time.Time, blurb string) *ComicAdded {
	return &ComicAdded{
		Id:            comicId,
		SeriesTitle:   seriesTitle,
		Title:         bookTitle,
		PublishedDate: publishedDate,
		WrittenBy:     writtenBy,
		ArtBy:         artBy,
		Blurb:         blurb,
		Pages:         pages,
	}
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
	return fmt.Sprintf("%T:{id: %v, "+
		"SeriesTitle: \"%v\", "+
		"BookTitle: \"%v\", "+
		"PublishedDate: %v, "+
		"WrittenBy: %#v, "+
		"ArtBy: %#v, "+
		"Blurb: %#v, "+
		"Pages: %v}",
		ca, ca.Id, ca.SeriesTitle, ca.Title, ca.PublishedDate, ca.WrittenBy, ca.ArtBy, ca.Blurb, ca.Pages)
}
