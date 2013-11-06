package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

func addComic(newId uuid.UUID, seriesTitle, bookTitle string, pages map[string]string, eventstorer EventStorer, filestorer FileStorer) error {
	series, err := domain.NewSeriesTitle(seriesTitle)
	if err != nil {
		return err
	}

	title, err := domain.NewBookTitle(bookTitle)
	if err != nil {
		return err
	}

	pagenames := getPageFilenames(pages)
	event := domain.NewComicAdded(domain.NewComicId(newId), series, title, pagenames)
	eventstorer.AddEvent(event)                // TODO: Deal with errors from the eventstorer
	filestorer.Store(event.Id.String(), pages) // TODO: Deal with errors from the filestorer

	return nil
}

func getPageFilenames(pages map[string]string) []string {
	names := make([]string, 0, len(pages))
	for key, _ := range pages {
		names = append(names, key)
	}
	return names
}
