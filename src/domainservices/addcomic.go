package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"errors"
	"log"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string, eventstorer EventStorer, filestorer FileStorer, vs frontPageViewStore) error {
	series, err := domain.NewSeriesTitle(seriesTitle)
	if err != nil {
		return err
	}

	title, err := domain.NewBookTitle(bookTitle)
	if err != nil {
		return err
	}

	if len(pages) == 0 {
		return errors.New("At least one page is required")
	}

	if len(pageSources) == 0 {
		return errors.New("At least one page source is required")
	}

	event := domain.NewComicAdded(domain.NewComicId(newId), series, title, pages)

	log.Printf("Storing new comic book %v\n", event)

	eventstorer.AddEvent(event)                             // TODO: Deal with errors from the eventstorer
	filestorer.Store(event.Id.String(), pages, pageSources) // TODO: Deal with errors from the filestorer
	saveFrontPage(vs, event)

	return nil
}

func saveFrontPage(vs frontPageViewStore, event *domain.ComicAdded) {
	frontPage := vs.Get()

	for i, s := range frontPage.Series {
		if s.Title == event.SeriesTitle.String() {
			frontPage.Series[i].ImageKey = event.Id.String() + "/" + event.Pages[0]
			vs.Store(&frontPage)
			return
		}
	}

	newseries := []FrontPageViewSeries{
		FrontPageViewSeries{
			Title:    event.SeriesTitle.String(),
			ImageKey: event.Id.String() + "/" + event.Pages[0],
		},
	}

	frontPage.Series = append(newseries, frontPage.Series...)
	vs.Store(&frontPage)
}
