package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
)

type ComicDomain struct {
	AddComic         func(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string) error
	GetFrontPageView func() *FrontPageView
}

func NewComicDomain(es EventStorer, fs FileStorer, vs ViewGetStorer) ComicDomain {
	frontPageVs := newFrontPageViewStore(vs)

	return ComicDomain{
		AddComic: func(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string) error {
			return addComic(newId, seriesTitle, bookTitle, pages, pageSources, es, fs, frontPageVs)
		},
		GetFrontPageView: func() *FrontPageView {
			fp := frontPageVs.Get()
			return &fp
		},
	}
}
