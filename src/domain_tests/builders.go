package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

func NewComicAdded(id uuid.UUID, seriesTitle, bookTitle string) *domain.ComicAdded {
	series, _ := domain.NewSeriesTitle(seriesTitle)
	book, _ := domain.NewBookTitle(bookTitle)
	comicId := domain.NewComicId(id)
	return domain.NewComicAdded(comicId, series, book)
}
