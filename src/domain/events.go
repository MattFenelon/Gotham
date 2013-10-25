package domain

type ComicAdded struct {
	Id          string // TODO: Create an identifier type
	SeriesTitle string
	BookTitle   string
}

func NewComicAdded(comicId string, seriesTitle, bookTitle string) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, BookTitle: bookTitle}
}
