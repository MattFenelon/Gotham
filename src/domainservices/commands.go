package domainservices

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
)

type CreateComicCommand struct {
	comicId     uuid.UUID
	seriesTitle domain.TrimmedString
	bookTitle   string
}

func NewCreateComicCommand(comicId uuid.UUID, seriesTitle, bookTitle string) *CreateComicCommand {
	// TODO: Validation
	// Trim string
	return &CreateComicCommand{comicId: comicId, seriesTitle: domain.ToTrimmedString(seriesTitle), bookTitle: bookTitle}
}
