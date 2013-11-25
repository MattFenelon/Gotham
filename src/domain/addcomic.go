package domain

import (
	"code.google.com/p/go-uuid/uuid"
	"domain/model"
	"errors"
	"log"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string, eventstorer EventStorer, filestorer FileStorer, fpVs *frontPageViewStore, comicVs *comicViewStore) error {
	series, err := model.NewSeriesTitle(seriesTitle)
	if err != nil {
		return err
	}

	title, err := model.NewBookTitle(bookTitle)
	if err != nil {
		return err
	}

	if len(pages) == 0 {
		return errors.New("At least one page is required")
	}

	if len(pageSources) == 0 {
		return errors.New("At least one page source is required")
	}

	event := model.NewComicAdded(model.NewComicId(newId), series, title, pages)

	log.Printf("Storing new comic book %v\n", event)

	eventstorer.AddEvent(event)                             // TODO: Deal with errors from the eventstorer
	filestorer.Store(event.Id.String(), pages, pageSources) // TODO: Deal with errors from the filestorer
	comicVs.Store(event)
	fpVs.Store(event)

	return nil
}
