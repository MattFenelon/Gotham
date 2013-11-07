package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"log"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, pages []string, pageSources []string, eventstorer EventStorer, filestorer FileStorer) error {
	series, err := domain.NewSeriesTitle(seriesTitle)
	if err != nil {
		return err
	}

	title, err := domain.NewBookTitle(bookTitle)
	if err != nil {
		return err
	}

	event := domain.NewComicAdded(domain.NewComicId(newId), series, title, pages)

	log.Printf("Storing new comic book %v\n", event)

	eventstorer.AddEvent(event)                             // TODO: Deal with errors from the eventstorer
	filestorer.Store(event.Id.String(), pages, pageSources) // TODO: Deal with errors from the filestorer

	return nil
}

func getPageFilenames(pages map[string]string) []string {
	names := make([]string, 0, len(pages))
	for key, _ := range pages {
		names = append(names, key)
	}
	return names
}
