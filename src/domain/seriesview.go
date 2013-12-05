package domain

import (
	"domain/model"
	"time"
)

type SeriesView struct {
	Title string
	Books []SeriesViewBookView
}

type SeriesViewBookView struct {
	Id            string
	Title         string
	PublishedDate time.Time
	WrittenBy     []string
	ArtBy         []string
	CoverBy       string
	Blurb         string
	ImageKey      string
}

type seriesViewStore struct {
	vs ViewGetStorer
}

func newSeriesViewStore(vs ViewGetStorer) *seriesViewStore {
	return &seriesViewStore{vs}
}

func (vs *seriesViewStore) Store(comicAdded *model.ComicAdded) {
	view := SeriesView{Books: []SeriesViewBookView{}}
	vs.vs.Get("SeriesView:"+comicAdded.SeriesTitle.String(), &view)

	bookView := SeriesViewBookView{
		Id:            comicAdded.Id.String(),
		Title:         comicAdded.Title.String(),
		PublishedDate: comicAdded.PublishedDate,
		WrittenBy:     comicAdded.WrittenBy,
		ArtBy:         comicAdded.ArtBy,
		Blurb:         comicAdded.Blurb,
		ImageKey:      comicAdded.Id.String() + "/" + comicAdded.Pages[0],
	}

	view.Title = comicAdded.SeriesTitle.String()
	view.Books = append([]SeriesViewBookView{bookView}, view.Books...)

	vs.vs.Store("SeriesView:"+comicAdded.SeriesTitle.String(), &view)
}

func (vs *seriesViewStore) Get(seriesTitle string) *SeriesView {
	var view SeriesView
	if err := vs.vs.Get("SeriesView:"+seriesTitle, &view); err != nil {
		return nil
	}
	return &view
}
