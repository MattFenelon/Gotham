package domain

import (
	"code.google.com/p/go-uuid/uuid"
	"domain/model"
	"errors"
	"log"
	"time"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, pages, pageSources, writtenBy, artBy []string, publishedDate time.Time, blurb string, eventstorer EventStorer, filestorer FileStorer, fpVs *frontPageViewStore, comicVs *comicViewStore, seriesVs *seriesViewStore) error {
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

	event := model.NewComicAdded(
		model.NewComicId(newId),
		series,
		title,
		writtenBy,
		artBy,
		pages,
		publishedDate,
		blurb)

	log.Printf("Storing new comic book %v\n", event)

	eventstorer.AddEvent(event)                             // TODO: Deal with errors from the eventstorer
	filestorer.Store(event.Id.String(), pages, pageSources) // TODO: Deal with errors from the filestorer
	comicVs.Store(event)
	fpVs.Store(event)
	seriesVs.Store(event)

	return nil
}
