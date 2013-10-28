package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
)

type ComicDomain struct {
	AddComic func(newId uuid.UUID, seriesTitle, bookTitle string)
}

func NewComicDomain(storer EventStorer) ComicDomain {
	return ComicDomain{
		AddComic: func(newId uuid.UUID, seriesTitle, bookTitle string) {
			addComic(newId, seriesTitle, bookTitle, storer)
		}}
}
