package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
)

type ComicDomain struct {
	AddComic func(newId uuid.UUID, seriesTitle, bookTitle string) error
}

func NewComicDomain(storer EventStorer) ComicDomain {
	return ComicDomain{
		AddComic: func(newId uuid.UUID, seriesTitle, bookTitle string) error {
			return addComic(newId, seriesTitle, bookTitle, storer)
		}}
}
