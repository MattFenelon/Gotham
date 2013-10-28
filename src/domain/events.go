package domain

type ComicAdded struct {
	Id          string // TODO: Create an identifier type
	SeriesTitle seriesTitle
	BookTitle   bookTitle
}

func NewComicAdded(comicId string, seriesTitle seriesTitle, bookTitle bookTitle) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, BookTitle: bookTitle}
}
