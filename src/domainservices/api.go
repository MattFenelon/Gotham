package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
)

type ComicDomain struct {
	AddComic func(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string) error
}

func NewComicDomain(eventStorer EventStorer, fileStorer FileStorer) ComicDomain {
	return ComicDomain{
		AddComic: func(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string) error {
			return addComic(newId, seriesTitle, bookTitle, pages, pageSources, eventStorer, fileStorer)
		}}
}
