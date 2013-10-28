package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, store EventStorer) {
	event := domain.NewComicAdded(newId.String(), seriesTitle, bookTitle)
	store.AddEvent(event)
}
