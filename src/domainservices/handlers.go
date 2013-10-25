package domainservices

import (
	"domain"
)

func AddComic(command *CreateComicCommand, store EventStorer) {
	event := domain.NewComicAdded(command.comicId.String(), command.seriesTitle, command.bookTitle)
	store.AddEvent(event)
}
