package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, eventstorer EventStorer, filestorer FileStorer) error {
	series, err := domain.NewSeriesTitle(seriesTitle)
	if err != nil {
		return err
	}

	title, err := domain.NewBookTitle(bookTitle)
	if err != nil {
		return err
	}

	event := domain.NewComicAdded(domain.NewComicId(newId), series, title, pages)
	eventstorer.AddEvent(event)                // TODO: Deal with errors from the eventstorer
	filestorer.Store(event.Id.String(), pages) // TODO: Deal with errors from the filestorer

	return nil
}
