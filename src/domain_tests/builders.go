package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain/model"
	"time"
)

func NewComicAdded(id uuid.UUID, seriesTitle, bookTitle string, writtenBy, artBy, pages []string, publishedDate time.Time, blurb string) *model.ComicAdded {
	series, _ := model.NewSeriesTitle(seriesTitle)
	book, _ := model.NewBookTitle(bookTitle)
	comicId := model.NewComicId(id)
	return model.NewComicAdded(comicId, series, book, writtenBy, artBy, pages, publishedDate, blurb)
}
