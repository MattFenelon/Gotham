package domain

type ComicAdded struct {
	Id          string // TODO: Create an identifier type
	SeriesTitle trimmedString
	BookTitle   trimmedString
}

func NewComicAdded(comicId string, seriesTitle, bookTitle trimmedString) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, BookTitle: bookTitle}
}
