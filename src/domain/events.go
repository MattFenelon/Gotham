package domain

type ComicAdded struct {
	Id          comicId // TODO: Create an identifier type
	SeriesTitle seriesTitle
	BookTitle   bookTitle
}

func NewComicAdded(comicId comicId, seriesTitle seriesTitle, bookTitle bookTitle) *ComicAdded {
	return &ComicAdded{Id: comicId, SeriesTitle: seriesTitle, BookTitle: bookTitle}
}
