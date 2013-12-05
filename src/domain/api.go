package domain

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

type ComicDomain struct {
	AddComic         func(newId uuid.UUID, seriesTitle, bookTitle string, pages, pageSources, writtenBy, artBy []string, publishedDate time.Time, blurb string) error
	GetFrontPageView func() *FrontPageView
	GetComicView     func(id uuid.UUID) *ComicView
	GetSeriesView    func(seriesTitle string) *SeriesView
}

func NewComicDomain(es EventStorer, fs FileStorer, vs ViewGetStorer) ComicDomain {
	frontPageVs := newFrontPageViewStore(vs)
	comicVs := newComicViewStore(vs)
	seriesVs := newSeriesViewStore(vs)

	return ComicDomain{
		AddComic: func(newId uuid.UUID, seriesTitle, bookTitle string, pages, pageSources, writtenBy, artBy []string, publishedDate time.Time, blurb string) error {
			return addComic(newId, seriesTitle, bookTitle, pages, pageSources, writtenBy, artBy, publishedDate, blurb, es, fs, frontPageVs, comicVs, seriesVs)
		},
		GetFrontPageView: func() *FrontPageView {
			fp := frontPageVs.Get()
			return &fp
		},
		GetComicView: func(id uuid.UUID) *ComicView {
			return comicVs.Get(id)
		},
		GetSeriesView: func(seriesTitle string) *SeriesView {
			return seriesVs.Get(seriesTitle)
		},
	}
}
