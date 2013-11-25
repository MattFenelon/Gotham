package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain/model"
)

func NewComicAdded(id uuid.UUID, seriesTitle, bookTitle string, pages []string) *model.ComicAdded {
	series, _ := model.NewSeriesTitle(seriesTitle)
	book, _ := model.NewBookTitle(bookTitle)
	comicId := model.NewComicId(id)
	return model.NewComicAdded(comicId, series, book, pages)
}
