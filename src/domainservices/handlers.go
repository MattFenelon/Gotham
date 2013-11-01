package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, store EventStorer) error {
	series, err := domain.NewSeriesTitle(seriesTitle)
	if err != nil {
		return err
	}

	title, err := domain.NewBookTitle(bookTitle)
	if err != nil {
		return err
	}

	event := domain.NewComicAdded(domain.NewComicId(newId), series, title)
	store.AddEvent(event) // TODO: Deal with errors from the store

	return nil
}
