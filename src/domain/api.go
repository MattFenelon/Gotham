package domain

import (
	"code.google.com/p/go-uuid/uuid"
)

type ComicDomain struct {
	AddComic         func(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string) error
	GetFrontPageView func() *FrontPageView
	GetComicView     func(id uuid.UUID) *ComicView
}

func NewComicDomain(es EventStorer, fs FileStorer, vs ViewGetStorer) ComicDomain {
	frontPageVs := newFrontPageViewStore(vs)
	comicVs := newComicViewStore(vs)

	return ComicDomain{
		AddComic: func(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string) error {
			return addComic(newId, seriesTitle, bookTitle, pages, pageSources, es, fs, frontPageVs, comicVs)
		},
		GetFrontPageView: func() *FrontPageView {
			fp := frontPageVs.Get()
			return &fp
		},
		GetComicView: func(id uuid.UUID) *ComicView {
			return comicVs.Get(id)
		},
	}
}
