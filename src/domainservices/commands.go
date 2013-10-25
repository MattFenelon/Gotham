package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
)

type CreateComicCommand struct {
	comicId     uuid.UUID
	seriesTitle string
	bookTitle   string
}

func NewCreateComicCommand(comicId uuid.UUID, seriesTitle, bookTitle string) *CreateComicCommand {
	return &CreateComicCommand{comicId: comicId, seriesTitle: seriesTitle, bookTitle: bookTitle}
}
