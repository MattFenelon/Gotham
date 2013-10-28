package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, store EventStorer) {
	trimmedSeries := domain.NewTrimmedString(seriesTitle)
	trimmedBook := domain.NewTrimmedString(bookTitle)

	event := domain.NewComicAdded(newId.String(), trimmedSeries, trimmedBook)
	store.AddEvent(event)
}
